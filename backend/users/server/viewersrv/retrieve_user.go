package viewersrv

import (
	"context"

	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/athome/backend/users/server"
	"github.com/athomecomar/athome/pb/pbusers"
	"github.com/jmoiron/sqlx"
)

func (s *Server) RetrieveUser(ctx context.Context, in *pbusers.RetrieveUserRequest) (*pbusers.User, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	user, err := server.FindUser(ctx, db, in.GetUserId())
	if err != nil {
		return nil, err
	}

	return s.retrieveUser(ctx, db, user)
}

func (s *Server) retrieveUser(ctx context.Context, db *sqlx.DB, user *ent.User) (*pbusers.User, error) {
	return user.ToPb(), nil
}
