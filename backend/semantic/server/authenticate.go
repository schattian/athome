package server

import (
	"context"

	"github.com/athomecomar/athome/backend/semantic/pb/pbauth"
	"github.com/athomecomar/athome/backend/semantic/pb/pbproducts"
	"github.com/athomecomar/athome/backend/semantic/semanticconf"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func ConnAuth(ctx context.Context) (pbauth.AuthClient, func() error, error) {
	conn, err := grpc.Dial(semanticconf.GetAUTH_ADDR(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, nil, status.Errorf(xerrors.Internal, "grpc.Dial: %v at %v", err, semanticconf.GetAUTH_ADDR())
	}
	c := pbauth.NewAuthClient(conn)
	return c, conn.Close, nil
}

func GetUserFromAccessToken(ctx context.Context, db *sqlx.DB, access string) (uint64, error) {
	c, closer, err := ConnAuth(ctx)
	if err != nil {
		return 0, err
	}
	defer closer()

	resp, err := c.RetrieveAuthentication(ctx, &pbauth.RetrieveAuthenticationRequest{AccessToken: access})
	if err != nil {
		return 0, err
	}

	return resp.GetUserId(), nil
}

func AuthorizeThroughEntity(ctx context.Context, access string, entityId uint64, entityTable string) (userId uint64, err error) {
	type authorizationFunc func(ctx context.Context, access string, entityId uint64) (userId uint64, err error)

	authFns := map[string]authorizationFunc{
		"drafts": authorizeProductsDrafts,
	}
	authFn, ok := authFns[entityTable]
	if !ok {
		return 0, status.Error(xerrors.InvalidArgument, "invalid entity table given")
	}

	return authFn(ctx, access, entityId)
}

func authorizeProductsDrafts(ctx context.Context, access string, entityId uint64) (uint64, error) {
	conn, err := grpc.Dial(semanticconf.GetPRODUCTS_ADDR(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return 0, status.Errorf(xerrors.Internal, "grpc.Dial: %v at %v", err, semanticconf.GetAUTH_ADDR())
	}
	defer conn.Close()
	c := pbproducts.NewCreatorClient(conn)
	draft, err := c.FetchDraft(ctx, &pbproducts.FetchDraftRequest{AccessToken: access})
	if err != nil {
		return 0, err
	}
	if draft.GetDraftId() != entityId {
		return 0, status.Errorf(xerrors.PermissionDenied, "your draft is %d, while trying to edit %d", draft.GetDraftId(), entityId)
	}

	return draft.GetUserId(), nil
}
