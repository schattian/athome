package server

import (
	"context"
	"strings"

	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/athome/backend/users/pb/pbuser"
	"github.com/athomecomar/storeql"
	_ "github.com/lib/pq"
	"github.com/omeid/pgerror"
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
	pqErr := storeql.InsertIntoDB(ctx, db, user)
	if pqErr.Is(pgerror.UniqueViolation) {
		return nil, errors.New("Ya existe un usuario con esa combinación de rol y email")
	}
	if err != nil {
		return nil, errors.Wrap(err, "storeql.InsertIntoDB")
	}
	return &pbuser.SignUpResponse{User: userToSignInUser(user)}, nil
}

func passwordHash(pwd string) (string, error) {
	err := field.Password(pwd).Validate()
	if err != nil {
		return "", errors.Wrap(err, "password.Validate")
	}
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
