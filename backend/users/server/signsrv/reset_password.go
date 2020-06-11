package signsrv

import (
	"context"
	"database/sql"

	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/athome/backend/users/server"
	"github.com/athomecomar/athome/backend/users/userconf"
	"github.com/athomecomar/athome/pb/pbusers"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ResetPassword(ctx context.Context, in *pbusers.ResetPasswordRequest) (*emptypb.Empty, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "server.ConnDB: %v", err)
	}
	defer db.Close()

	return s.resetPassword(ctx, db, in)
}

func (s *Server) resetPassword(ctx context.Context, db *sqlx.DB, in *pbusers.ResetPasswordRequest) (*emptypb.Empty, error) {
	userId, err := handleJwt(in.GetToken(), userconf.GetFORGOT_JWT_SECRET)
	if err != nil {
		return nil, err
	}

	u := &ent.User{}
	row := db.QueryRowxContext(ctx, `SELECT * FROM users WHERE id=$1`, userId)
	err = row.StructScan(u)
	if err == sql.ErrNoRows {
		return nil, status.Errorf(xerrors.NotFound, "StructScan: %v", err)
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "StructScan: %v", err)
	}
	err = u.AssignPassword(in.GetPassword())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "AssignPassword: %v", err)
	}
	err = storeql.UpdateIntoDB(ctx, db, u)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "storeql.UpdateIntoDB: %v", err)
	}

	return &emptypb.Empty{}, nil
}
