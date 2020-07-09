package srvshippings

import (
	"context"

	"github.com/athomecomar/athome/backend/checkout/ent/shipping"
	"github.com/athomecomar/athome/backend/checkout/ent/sm"
	"github.com/athomecomar/athome/backend/checkout/server"
	"github.com/athomecomar/athome/pb/pbagreement"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) AgreeShippingDispatch(ctx context.Context, in *pbcheckout.AgreeShippingDispatchRequest) (*emptypb.Empty, error) {
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
	sh, err := shipping.FindShipping(ctx, db, in.GetShippingId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindShipping: %v", err)
	}

	agree, agreeCloser, err := pbutil.ConnAgreement(ctx)
	if err != nil {
		return nil, err
	}

	defer agreeCloser()
	return s.agreeShippingDispatch(ctx, db, agree, in.GetAgreeementToken(), sh, uid)
}

func (s *Server) agreeShippingDispatch(
	ctx context.Context,
	db *sqlx.DB,
	agree pbagreement.AgreementClient,
	token uint64,
	sh *shipping.Shipping,
	uid uint64,
) (*emptypb.Empty, error) {
	if uid == sh.UserId {
		return nil, status.Error(xerrors.PermissionDenied, "you cant agree with yourself")
	}
	if sh.GetMerchantId() != uid {
		return nil, status.Errorf(xerrors.PermissionDenied, "you cant agree shipment dispatches if you aren't the merchant")
	}
	_, err := agree.Verify(ctx, &pbagreement.VerifyRequest{AgreedUserId: sh.UserId, AgreementToken: token})
	if err != nil {
		return nil, err
	}
	err = server.MustPrevState(ctx, db, sh, sm.ShippingTaken, sh.UserId)
	if err != nil {
		return nil, err
	}
	err = server.ChangeState(ctx, db, sm.Next, sh, sh.UserId)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
