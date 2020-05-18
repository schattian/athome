package server

import (
	"context"

	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc/status"

	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/athomecomar/athome/pb/go/pbuser"
	_ "github.com/lib/pq"
)

func (s *Server) AlreadyExists(ctx context.Context, in *pbuser.AlreadyExistsRequest) (*pbuser.AlreadyExistsResponse, error) {
	db, err := connDB()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "connDB: %v", err)
	}
	email := in.GetEmail()
	if email == "" {
		return nil, status.Error(xerrors.InvalidArgument, "no email given")
	}
	role := in.GetRole()
	if role == "" {
		return nil, status.Error(xerrors.InvalidArgument, "no role given")
	}
	// validations'll be fully replaced by input validations using go-grpc-validate (see .proto)
	if err := field.Role(role).Validate(); err != nil {
		return nil, status.Error(xerrors.InvalidArgument, "invalid role given")
	}
	if err := field.Email(email).Validate(); err != nil {
		return nil, status.Error(xerrors.InvalidArgument, "invalid email given")
	}
	rows, err := db.QueryxContext(ctx, `SELECT COUNT(*) FROM users WHERE email=$1 AND role=$2 LIMIT 1`, email, role)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "QueryxContext: %v", err)
	}
	defer rows.Close()

	var count int64
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return nil, status.Errorf(xerrors.Internal, "rows.Scan: %v", err)
		}
	}
	return &pbuser.AlreadyExistsResponse{Exists: count > 0}, nil
}
