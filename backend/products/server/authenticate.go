package server

import (
	"context"

	"github.com/athomecomar/athome/backend/products/ent"
	"github.com/athomecomar/athome/backend/products/pb/pbauth"
	"github.com/athomecomar/athome/backend/products/productconf"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func GetUserFromAccessToken(ctx context.Context, c pbauth.AuthClient, access string) (uint64, error) {
	resp, err := c.RetrieveAuthentication(ctx, &pbauth.RetrieveAuthenticationRequest{AccessToken: access})
	if err != nil {
		return 0, err
	}

	return resp.GetUserId(), nil
}

func ConnAuth(ctx context.Context) (pbauth.AuthClient, func() error, error) {
	conn, err := grpc.Dial(productconf.GetAUTH_ADDR(), grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		return nil, nil, status.Errorf(xerrors.Internal, "grpc.Dial: %v at %v", err, productconf.GetAUTH_ADDR())
	}
	c := pbauth.NewAuthClient(conn)
	return c, conn.Close, nil
}

func RetrieveLatestDraft(ctx context.Context, db *sqlx.DB, accessToken string) (*ent.Draft, error) {
	c, closer, err := ConnAuth(ctx)
	if err != nil {
		return nil, err
	}
	draft, err := retrieveLatestDraft(ctx, db, c, closer, accessToken)
	if err != nil {
		return nil, err
	}

	return draft, nil
}

func retrieveLatestDraft(ctx context.Context, db *sqlx.DB, auth pbauth.AuthClient, authCloser func() error, accessToken string) (*ent.Draft, error) {
	defer authCloser()
	userId, err := GetUserFromAccessToken(ctx, auth, accessToken)
	if err != nil {
		return nil, err
	}

	draft, err := ent.FindOrCreateDraft(ctx, db, userId)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindOrCreateDraft: %v", err)
	}
	if draft.Id > 0 {
		return draft, nil
	}
	err = storeql.InsertIntoDB(ctx, db, draft)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "InsertIntoDB: %v", err)
	}

	return draft, nil
}
