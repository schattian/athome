package server

import (
	"context"

	"github.com/athomecomar/athome/users/ent"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"

	"github.com/athomecomar/athome/users/pb/pbuser"
	"github.com/pkg/errors"
)

func (s *Server) SignIn(ctx context.Context, in *pbuser.SignInRequest) (*pbuser.SignInResponse, error) {
	db, err := connDB()
	if err != nil {
		return nil, errors.Wrap(err, "connDB")
	}
	email := in.GetEmail()
	if email == "" {
		return nil, errors.Wrap(err, "no email given")
	}
	rows, err := db.QueryxContext(ctx, `SELECT * FROM users WHERE email=$1 limit 3`, email)
	if err != nil {
		return nil, errors.Wrap(err, "QueryxContext")
	}
	var users []*pbuser.SignInUser
	defer rows.Close()
	for rows.Next() {
		user := &ent.User{}
		err = rows.StructScan(user)
		if err != nil {
			return nil, errors.Wrap(err, "rows.StructScan")
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(in.GetPassword()))
		if err != nil {
			continue
		}

		users = append(users, userToSignInUser(user))
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.Wrap(err, "rows.Err")
	}
	return &pbuser.SignInResponse{Users: users}, nil
}

func userToSignInUser(user *ent.User) *pbuser.SignInUser {
	return &pbuser.SignInUser{Token: user.PasswordHash, Role: string(user.Role)}
}
