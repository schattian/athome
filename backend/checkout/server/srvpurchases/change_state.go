package srvpurchases

import (
	"context"

	"github.com/athomecomar/athome/backend/checkout/ent/order"
	"github.com/athomecomar/athome/backend/checkout/ent/sm"
	"github.com/athomecomar/athome/backend/checkout/server"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) Cancel(ctx context.Context, in *pbcheckout.UpdateOrderStateRequest) (*pbcheckout.RetrievePurchaseResponse, error) {
	return s.ChangeState(ctx, in, sm.Cancel)
}

func (s *Server) Prev(ctx context.Context, in *pbcheckout.UpdateOrderStateRequest) (*pbcheckout.RetrievePurchaseResponse, error) {
	return s.ChangeState(ctx, in, sm.Prev)
}

func (s *Server) Next(ctx context.Context, in *pbcheckout.UpdateOrderStateRequest) (*pbcheckout.RetrievePurchaseResponse, error) {
	return s.ChangeState(ctx, in, sm.Next)
}

func (s *Server) ChangeState(ctx context.Context, in *pbcheckout.UpdateOrderStateRequest,
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

	return s.changeState(ctx, db, stateChanger, prods, o)
}

func (s *Server) changeState(
	ctx context.Context,
	db *sqlx.DB,
	stateChanger sm.StateChanger,
	prods pbproducts.ViewerClient,
	o *order.Purchase,
) (*pbcheckout.RetrievePurchaseResponse, error) {
	sc, err := order.LatestStateChange(ctx, db, o)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "LatestStateChange")
	}
	state, err := stateChanger(o.StateMachine(), sc.GetState())
	if err != nil {
		return nil, status.Errorf(xerrors.InvalidArgument, "sm stateChanger: %v", err)
	}
	err = o.ValidateStateChange(ctx, db, state)
	if err != nil {
		return nil, status.Errorf(xerrors.InvalidArgument, "ValidateStateChange: %v", err)
	}
	sc, err = order.NewPurchaseStateChange(ctx, o.Id, state.Name)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "NewPurchaseStateChange: %v", err)
	}
	err = storeql.InsertIntoDB(ctx, db, sc)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "storeql.InsertIntoDB: %v", err)
	}

	return s.retrievePurchase(ctx, db, prods, o)
}
