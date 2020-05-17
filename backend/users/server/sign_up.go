package server

import (
	"context"
	"strings"

	"github.com/athomecomar/athome/users/ent"
	"github.com/athomecomar/athome/users/ent/field"
	"github.com/athomecomar/athome/users/pb/pbuser"
	"github.com/athomecomar/storeql"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) SignUp(ctx context.Context, in *pbuser.SignUpRequest) (*pbuser.SignUpResponse, error) {
	hash, err := passwordHash(in.GetPassword())
	if err != nil {
		return nil, errors.Wrap(err, "passwordHash")
	}
	db, err := connDB()
	if err != nil {
		return nil, errors.Wrap(err, "connDB")
	}
	user, err := signUpUserToUser(in, hash)
	if err != nil {
		return nil, errors.Wrap(err, "signUpUserToUser")
	}

	err = storeql.InsertIntoDB(ctx, db, user)
	if err != nil {
		return nil, errors.Wrap(err, "storeql.InsertIntoDB")
	}

	return &pbuser.SignUpResponse{}, nil
}

func passwordHash(pwd string) (string, error) {
	ph, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.Wrap(err, "bcrypt.GenerateFromPassword")
	}
	return string(ph), nil
}

func signUpUserToUser(in *pbuser.SignUpRequest, pwdHash string) (*ent.User, error) {
	user := &ent.User{
		Email:        field.Email(strings.ToLower(string(in.GetEmail()))),
		PasswordHash: pwdHash,
		Role:         field.Role(in.GetRole()),
		Name:         field.Name(in.GetName()),
		Surname:      field.Surname(in.GetSurname()),
	}
	if err := user.Validate(); err != nil {
		return nil, errors.Wrap(err, "user.Validate")
	}
	return user, nil
}
