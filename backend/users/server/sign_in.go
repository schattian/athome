package server

import (
	"context"

	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"

	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/athome/backend/users/pbuser"
	"github.com/athomecomar/athome/backend/users/userconf"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) SignIn(ctx context.Context, in *pbuser.SignInRequest) (*pbuser.SignInResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	db, err := connDB()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "connDB: %v", err)
	}
	defer db.Close()
	return s.signIn(ctx, db, in)
}

func (s *Server) signIn(ctx context.Context, db *sqlx.DB, in *pbuser.SignInRequest) (*pbuser.SignInResponse, error) {
	rows, err := db.QueryxContext(ctx, `SELECT * FROM users WHERE email=$1 limit 3`, in.GetEmail())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "QueryxContext: %v", err)
	}
	var users []*pbuser.SignInUser
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
	return &pbuser.SignInResponse{
		Users:          users,
		SignTokenExpNs: uint64(userconf.GetSIGN_JWT_EXP().Nanoseconds()),
	}, nil
}

func userToSignInUser(user *ent.User) (*pbuser.SignInUser, error) {
	token, err := createSignToken(user.Id)
	if err != nil {
		return nil, errors.Wrap(err, "CreateSignToken")
	}
	return &pbuser.SignInUser{
		Id:        user.Id,
		SignToken: token,
		Email:     string(user.Email),
		Role:      string(user.Role),
		Name:      string(user.Name),
		Surname:   string(user.Surname),
	}, nil
}
