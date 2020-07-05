package srvshippings

import (
	"context"

	"github.com/athomecomar/athome/backend/checkout/ent/shipping"
	"github.com/athomecomar/athome/backend/checkout/ent/sm"
	"github.com/athomecomar/athome/backend/checkout/server"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) Cancel(ctx context.Context, in *pbcheckout.UpdateStateRequest) (*pbcheckout.RetrieveShippingResponse, error) {
	return s.ChangeState(ctx, in, sm.Cancel)
}

func (s *Server) Prev(ctx context.Context, in *pbcheckout.UpdateStateRequest) (*pbcheckout.RetrieveShippingResponse, error) {
	return s.ChangeState(ctx, in, sm.Prev)
}

func (s *Server) Next(ctx context.Context, in *pbcheckout.UpdateStateRequest) (*pbcheckout.RetrieveShippingResponse, error) {
	return s.ChangeState(ctx, in, sm.Next)
}

func (s *Server) ChangeState(ctx context.Context, in *pbcheckout.UpdateStateRequest,
	stateChanger sm.StateChanger,
) (*pbcheckout.RetrieveShippingResponse, error) {
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
	sh, err := shipping.FindShipping(ctx, db, in.GetEntityId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindShipping: %v", err)
	}

	return s.changeState(ctx, db, stateChanger, sh, uid)
}

func (s *Server) changeState(
	ctx context.Context,
	db *sqlx.DB,
	stateChanger sm.StateChanger,
	sh *shipping.Shipping,
	uid uint64,
) (*pbcheckout.RetrieveShippingResponse, error) {
	sc, err := sm.LatestStateChange(ctx, db, sh)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "LatestStateChange")
	}
	state, err := stateChanger(sh.StateMachine(), sc.GetState(), sh, uid)
	if err != nil {
		return nil, status.Errorf(xerrors.InvalidArgument, "sm stateChanger: %v", err)
	}
	err = sh.ValidateStateChange(ctx, db, state)
	if err != nil {
		return nil, status.Errorf(xerrors.InvalidArgument, "ValidateStateChange: %v", err)
	}
	sc, err = shipping.NewShippingStateChange(ctx, sh.Id, state.Name)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "NewShippingStateChange: %v", err)
	}
	err = storeql.InsertIntoDB(ctx, db, sc)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "storeql.InsertIntoDB: %v", err)
	}
	return &pbcheckout.RetrieveShippingResponse{}, nil
}
