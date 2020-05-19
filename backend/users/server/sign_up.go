package server

import (
	"context"
	"strings"

	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/athome/backend/users/pbuser"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	_ "github.com/lib/pq"
	"github.com/omeid/pgerror"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/status"
)

func (s *Server) SignUp(ctx context.Context, in *pbuser.SignUpRequest) (*pbuser.SignUpResponse, error) {
	pwd := in.GetPassword()
	err := field.Password(pwd).Validate()
	if err != nil {
		return nil, status.Errorf(xerrors.InvalidArgument, "password.Validate: %v", err)
	}

	hash, err := passwordHash(pwd)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "passwordHash: %v", err)
	}
	db, err := connDB()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "connDB: %v", err)
	}
	user, err := signUpUserToUser(in, hash)
	if err != nil {
		return nil, status.Errorf(xerrors.InvalidArgument, "signUpUserToUser: %v", err)
	}
	pqErr := storeql.InsertIntoDB(ctx, db, user)
	if pqErr.Is(pgerror.UniqueViolation) {
		return nil, status.Error(xerrors.AlreadyExists, "Ya existe un usuario con esa combinación de rol y email")
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "storeql.InsertIntoDB: %v", err)
	}
	return &pbuser.SignUpResponse{User: userToSignInUser(user)}, nil
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
