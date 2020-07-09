package srvpurchases

import (
	"context"

	"github.com/athomecomar/athome/backend/checkout/ent/sm"
	"github.com/athomecomar/athome/backend/checkout/server"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/athome/pb/pbutil"
)

func (s *Server) Cancel(ctx context.Context, in *pbcheckout.UpdateStateRequest) (*pbcheckout.RetrievePurchaseResponse, error) {
	return s.ChangeState(ctx, in, sm.Cancel)
}

func (s *Server) Prev(ctx context.Context, in *pbcheckout.UpdateStateRequest) (*pbcheckout.RetrievePurchaseResponse, error) {
	return s.ChangeState(ctx, in, sm.Prev)
}

func (s *Server) Next(ctx context.Context, in *pbcheckout.UpdateStateRequest) (*pbcheckout.RetrievePurchaseResponse, error) {
	return s.ChangeState(ctx, in, sm.Next)
}

func (s *Server) ChangeState(ctx context.Context, in *pbcheckout.UpdateStateRequest,
	stateChanger sm.StateChanger,
) (*pbcheckout.RetrievePurchaseResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	auth, authCloser, err := pbutil.ConnAuth(ctx)
	if err != nil {
		return nil, err
	}
	uid, err := pbutil.GetUserFromAccessToken(ctx, auth, in.GetAccessToken())
	if err != nil {
		return nil, err
	}
	err = authCloser()
	if err != nil {
		return nil, err
	}
	o, err := server.FindLatestPurchase(ctx, db, uid)
	if err != nil {
		return nil, err
	}
	prods, prodsCloser, err := pbutil.ConnProductsViewer(ctx)
	if err != nil {
		return nil, err
	}
	defer prodsCloser()

	err = server.ChangeState(ctx, db, stateChanger, o, uid)
	if err != nil {
		return nil, err
	}
	return s.retrievePurchase(ctx, db, prods, o)
}

// func (s *Server) changeState(
// 	ctx context.Context,
// 	db *sqlx.DB,
// 	stateChanger sm.StateChanger,
// 	o *purchase.Purchase,
// 	uid uint64,
// ) error {
// 	sc, err := sm.LatestStateChange(ctx, db, o)
// 	if err != nil {
// 		return status.Errorf(xerrors.Internal, "LatestStateChange")
// 	}
// 	state, err := stateChanger(o.StateMachine(), sc.GetState(o.StateMachine()), o, uid)
// 	if err != nil {
// 		return status.Errorf(xerrors.InvalidArgument, "sm stateChanger: %v", err)
// 	}
// 	err = o.ValidateStateChange(ctx, db, state)
// 	if err != nil {
// 		return status.Errorf(xerrors.InvalidArgument, "ValidateStateChange: %v", err)
// 	}
// 	sc, err = sm.NewStateChange(ctx, o.Id, state.Name, o)
// 	if err != nil {
// 		return status.Errorf(xerrors.Internal, "NewPurchaseStateChange: %v", err)
// 	}
// 	err = storeql.InsertIntoDB(ctx, db, sc)
// 	if err != nil {
// 		return status.Errorf(xerrors.Internal, "storeql.InsertIntoDB: %v", err)
// 	}
// 	return nil
// }
