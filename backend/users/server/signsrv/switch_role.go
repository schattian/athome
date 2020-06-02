package signsrv

import (
	"context"
	"database/sql"
	"time"

	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/athome/backend/users/internal/userjwt"
	"github.com/athomecomar/athome/backend/users/pb/pbauth"
	"github.com/athomecomar/athome/backend/users/pb/pbuser"
	"github.com/athomecomar/athome/backend/users/server"
	"github.com/athomecomar/athome/backend/users/userconf"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func (s *Server) SwitchRole(ctx context.Context, in *pbuser.SwitchRoleRequest) (*pbuser.SignResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	db, err := server.ConnDB()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "server.ConnDB: %v", err)
	}
	defer db.Close()

	conn, err := grpc.Dial(userconf.GetAUTH_ADDR(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "grpc.Dial: %v at %v", err, userconf.GetAUTH_ADDR())
	}
	defer conn.Close()
	c := pbauth.NewAuthClient(conn)

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	return s.switchRole(ctx, db, c, in)
}

func (s *Server) switchRole(ctx context.Context, db *sqlx.DB, c pbauth.AuthClient, in *pbuser.SwitchRoleRequest) (*pbuser.SignResponse, error) {
	oldUser, err := server.GetUserFromAccessToken(ctx, db, c, in.GetAccessToken())
	if err != nil {
		return nil, err
	}

	row := db.QueryRowxContext(ctx, `SELECT * FROM users WHERE email=$1 AND role=$2`, oldUser.Email, in.GetRole())
	user := &ent.User{}
	err = row.StructScan(user)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Errorf(xerrors.NotFound, "user StructScan: %v", err)
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "user StructScan: %v", err)
	}

	if user.PasswordHash != oldUser.PasswordHash {
		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(in.GetPassword()))
		if err != nil {
			return nil, status.Errorf(xerrors.Unauthenticated, "bcrypt.CompareHashAndPassword: %v", err)
		}
	}

	signToken, err := userjwt.CreateSignToken(user.Id)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "createSignToken: %v", err)
	}

	return s.sign(ctx, c, &pbuser.SignRequest{SignToken: signToken})
}
