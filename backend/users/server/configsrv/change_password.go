package configsrv

import (
	"context"
	"time"

	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/athomecomar/athome/backend/users/server"
	"github.com/athomecomar/athome/pb/pbauth"
	"github.com/athomecomar/athome/pb/pbconf"
	"github.com/athomecomar/athome/pb/pbusers"
	_ "github.com/lib/pq"
)

func (s *Server) ChangePassword(ctx context.Context, in *pbusers.ChangePasswordRequest) (*emptypb.Empty, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	db, err := server.ConnDB()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "server.ConnDB: %v", err)
	}
	defer db.Close()

	auth, authCloser, err := pbconf.ConnAuth(ctx)
	if err != nil {
		return nil, err
	}
	defer authCloser()

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	return s.changePassword(ctx, db, auth, in)
}

func (s *Server) changePassword(ctx context.Context, db *sqlx.DB, c pbauth.AuthClient, in *pbusers.ChangePasswordRequest) (*emptypb.Empty, error) {
	user, err := server.GetUserFromAccessToken(ctx, db, c, in.GetAccessToken())
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(in.GetOldPassword()))
	if err != nil {
		return nil, status.Errorf(xerrors.Unauthenticated, "bcrypt.CompareHashAndPassword: %v", err)
	}
	err = user.AssignPassword(in.GetNewPassword())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "AssignPassword: %v", err)
	}
	err = storeql.UpdateIntoDB(ctx, db, user)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "storeql.UpdateIntoDB: %v", err)
	}
	return &emptypb.Empty{}, nil
}
