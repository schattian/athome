package server

import (
	"context"

	"github.com/athomecomar/athome/backend/products/ent"
	"github.com/athomecomar/athome/pb/pbauth"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func RetrieveLatestDraft(ctx context.Context, db *sqlx.DB, accessToken string) (*ent.Draft, error) {
	c, closer, err := pbutil.ConnAuth(ctx)
	if err != nil {
		return nil, err
	}
	defer closer()
	draft, err := retrieveLatestDraft(ctx, db, c, accessToken)
	if err != nil {
		return nil, err
	}

	return draft, nil
}

func retrieveLatestDraft(ctx context.Context, db *sqlx.DB, auth pbauth.AuthClient, accessToken string) (*ent.Draft, error) {
	userId, err := pbutil.GetUserFromAccessToken(ctx, auth, accessToken)
	if err != nil {
		return nil, err
	}

	draft, err := ent.FindOrCreateDraft(ctx, db, userId)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindOrCreateDraft: %v", err)
	}

	return draft, nil
}
