package srvcards

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/athomecomar/athome/backend/checkout/ent/payment"
	"github.com/athomecomar/athome/backend/checkout/server"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) VerifyCVV(ctx context.Context, in *pbcheckout.VerifyCVVRequest) (*emptypb.Empty, error) {
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

	return s.verifyCVV(ctx, db, in, uid)
}

func (s *Server) verifyCVV(ctx context.Context, db *sqlx.DB, in *pbcheckout.VerifyCVVRequest, uid uint64) (*emptypb.Empty, error) {
	card, err := payment.FindCard(ctx, db, in.GetCardId(), uid)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Errorf(xerrors.NotFound, "FindCard: %v", err)
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindCard: %v", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(card.CVVHash), []byte(strconv.Itoa(int(in.GetCvv()))))
	if err != nil {
		return nil, status.Errorf(xerrors.PermissionDenied, "CompareHashAndPassword: %v", err)
	}
	return &emptypb.Empty{}, nil
}
