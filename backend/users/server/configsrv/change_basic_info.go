package configsrv

import (
	"context"
	"time"

	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/athome/backend/users/server"
	"github.com/athomecomar/athome/pb/pbauth"
	"github.com/athomecomar/athome/pb/pbusers"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ChangeBasicInfo(ctx context.Context, in *pbusers.ChangeBasicInfoRequest) (*emptypb.Empty, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	db, err := server.ConnDB()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "server.ConnDB: %v", err)
	}
	defer db.Close()

	auth, authCloser, err := pbutil.ConnAuth(ctx)
	if err != nil {
		return nil, err
	}
	defer authCloser()

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	return s.changeBasicInfo(ctx, db, auth, in)
}

func (s *Server) changeBasicInfo(ctx context.Context, db *sqlx.DB, c pbauth.AuthClient, in *pbusers.ChangeBasicInfoRequest) (*emptypb.Empty, error) {
	userId, err := pbutil.GetUserFromAccessToken(ctx, c, in.GetAccessToken())
	if err != nil {
		return nil, err
	}
	user, err := server.FindUser(ctx, db, userId)
	if err != nil {
		return nil, err
	}

	user.Name, user.Surname = field.Name(in.GetName()), field.Surname(in.GetSurname())
	err = storeql.UpdateIntoDB(ctx, db, user)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "storeql.UpdateIntoDB: %v", err)
	}
	return &emptypb.Empty{}, nil
}
