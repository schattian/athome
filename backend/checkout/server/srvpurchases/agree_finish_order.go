package srvpurchases

import (
	"context"

	"github.com/athomecomar/athome/backend/checkout/ent/order/purchase"
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

func (s *Server) AgreeFinishOrder(ctx context.Context, in *pbcheckout.AgreeFinishOrderRequest) (*emptypb.Empty, error) {
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
	p, err := purchase.FindPurchase(ctx, db, in.GetOrderId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindPurchase: %v", err)
	}
	agree, agreeCloser, err := pbutil.ConnAgreement(ctx)
	if err != nil {
		return nil, err
	}

	defer agreeCloser()
	return s.agreeFinishOrder(ctx, db, agree, in.GetAgreeementToken(), p, uid)
}

func (s *Server) agreeFinishOrder(
	ctx context.Context,
	db *sqlx.DB,
	agree pbagreement.AgreementClient,
	token uint64,
	p *purchase.Purchase,
	uid uint64,
) (*emptypb.Empty, error) {
	if uid == p.UserId {
		return nil, status.Error(xerrors.PermissionDenied, "you cant agree with yourself")
	}
	canView, err := p.CanView(ctx, db, uid)
	if err != nil || !canView {
		return nil, status.Errorf(xerrors.PermissionDenied, "you cant agree non owned orders: %v", err)
	}
	_, err = agree.Verify(ctx, &pbagreement.VerifyRequest{AgreedUserId: p.UserId, AgreementToken: token})
	if err != nil {
		return nil, err
	}

	if p.ShippingId != 0 {
		ship, err := p.Shipping(ctx, db)
		if err != nil {
			return nil, err
		}
		err = server.ChangeState(ctx, db, sm.Next, ship, p.UserId)
		if err != nil {
			return nil, err
		}
	}

	err = server.MustPrevState(ctx, db, p, sm.PurchaseFinished, p.UserId)
	if err != nil {
		return nil, err
	}
	err = server.ChangeState(ctx, db, sm.Next, p, p.UserId)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
