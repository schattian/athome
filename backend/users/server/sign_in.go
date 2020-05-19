package server

import (
	"context"

	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc/status"

	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/athome/backend/users/pbuser"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) SignIn(ctx context.Context, in *pbuser.SignInRequest) (*pbuser.SignInResponse, error) {
	db, err := connDB()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "connDB: %v", err)
	}
	email := in.GetEmail()
	if email == "" {
		return nil, status.Error(xerrors.InvalidArgument, "no email given")
	}
	rows, err := db.QueryxContext(ctx, `SELECT * FROM users WHERE email=$1 limit 3`, email)
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

		users = append(users, userToSignInUser(user))
	}
	err = rows.Err()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "rows.Err: %v", err)
	}
	return &pbuser.SignInResponse{Users: users}, nil
}

func userToSignInUser(user *ent.User) *pbuser.SignInUser {
	return &pbuser.SignInUser{
		Id:      user.Id,
		Token:   user.PasswordHash,
		Email:   string(user.Email),
		Role:    string(user.Role),
		Name:    string(user.Name),
		Surname: string(user.Surname),
	}
}
