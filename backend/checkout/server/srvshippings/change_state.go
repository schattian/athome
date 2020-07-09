package srvshippings

import (
	"context"

	"github.com/athomecomar/athome/backend/checkout/ent/shipping"
	"github.com/athomecomar/athome/backend/checkout/ent/sm"
	"github.com/athomecomar/athome/backend/checkout/server"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/xerrors"
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

	err = server.ChangeState(ctx, db, stateChanger, sh, uid)
	if err != nil {
		return nil, err
	}

	return &pbcheckout.RetrieveShippingResponse{ShippingId: sh.Id, Shipping: sh.ToPb()}, nil
}
