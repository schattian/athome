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
