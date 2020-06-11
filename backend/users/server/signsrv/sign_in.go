package signsrv

import (
	"context"

	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"

	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/athome/backend/users/internal/userjwt"
	"github.com/athomecomar/athome/backend/users/server"
	"github.com/athomecomar/athome/backend/users/userconf"
	"github.com/athomecomar/athome/pb/pbusers"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) SignIn(ctx context.Context, in *pbusers.SignInRequest) (*pbusers.SignInResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	db, err := server.ConnDB()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "server.ConnDB: %v", err)
	}
	defer db.Close()
	return s.signIn(ctx, db, in)
}

func (s *Server) signIn(ctx context.Context, db *sqlx.DB, in *pbusers.SignInRequest) (*pbusers.SignInResponse, error) {
	rows, err := db.QueryxContext(ctx, `SELECT * FROM users WHERE email=$1 limit 3`, in.GetEmail())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "QueryxContext: %v", err)
	}
	var users []*pbusers.SignInUser
	defer rows.Close()
	for rows.Next() {
		user := &ent.User{}
		err = rows.StructScan(user)
		if err != nil {
			return nil, status.Errorf(xerrors.Internal, "rows.StructScan: %v", err)
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(in.GetPassword()))
		if err != nil {
			continue
		}

		signedUser, err := userToSignInUser(user)
		if err != nil {
			return nil, status.Errorf(xerrors.Internal, "userToSignInUser9: %v", err)
		}

		users = append(users, signedUser)
	}

	err = rows.Err()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "rows.Err: %v", err)
	}
	return &pbusers.SignInResponse{
		Users:          users,
		SignTokenExpNs: uint64(userconf.GetSIGN_JWT_EXP().Nanoseconds()),
	}, nil
}

func userToSignInUser(user *ent.User) (*pbusers.SignInUser, error) {
	token, err := userjwt.CreateSignToken(user.Id)
	if err != nil {
		return nil, errors.Wrap(err, "CreateSignToken")
	}
	return &pbusers.SignInUser{
		Id:        user.Id,
		SignToken: token,
		User:      user.ToPb(),
	}, nil
}
