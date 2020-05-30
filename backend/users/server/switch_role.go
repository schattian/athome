package server

import (
	"context"
	"database/sql"
	"time"

	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/athome/backend/users/pb/pbauth"
	"github.com/athomecomar/athome/backend/users/pb/pbuser"
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

	db, err := connDB()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "connDB: %v", err)
	}
	defer db.Close()

	conn, err := grpc.Dial(userconf.GetAUTH_ADDR(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "grpc.Dial: %v at %v", err, userconf.GetAUTH_ADDR())
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	return s.switchRole(ctx, db, conn, in)
}

func (s *Server) switchRole(ctx context.Context, db *sqlx.DB, conn *grpc.ClientConn, in *pbuser.SwitchRoleRequest) (*pbuser.SignResponse, error) {
	c := pbauth.NewAuthClient(conn)
	userId, err := c.RetrieveAuthentication(ctx, &pbauth.RetrieveAuthenticationRequest{AccessToken: in.GetAccessToken()})
	if err != nil {
		return nil, err
	}

	oldUser := &ent.User{}
	row := db.QueryRowxContext(ctx, `SELECT * FROM users WHERE id=$1`, userId.GetUserId())
	err = row.StructScan(oldUser)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "oldUser row.StructScan: %v", err)
	}

	row = db.QueryRowxContext(ctx, `SELECT * FROM users WHERE email=$1 AND role=$2`, oldUser.Email, in.GetRole())
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

	signToken, err := createSignToken(user.Id)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "createSignToken: %v", err)
	}

	return s.sign(ctx, conn, &pbuser.SignRequest{SignToken: signToken})
}
