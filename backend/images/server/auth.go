package server

import (
	"context"

	"github.com/athomecomar/athome/pb/pbauth"
	"github.com/athomecomar/athome/pb/pbconf"
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc/status"
)

func GetUserFromAccessToken(ctx context.Context, c pbauth.AuthClient, access string) (uint64, error) {
	resp, err := c.RetrieveAuthentication(ctx, &pbauth.RetrieveAuthenticationRequest{AccessToken: access})
	if err != nil {
		return 0, err
	}

	return resp.GetUserId(), nil
}

func AuthorizeThroughEntity(ctx context.Context, access string, entityId uint64, entityTable string) (userId uint64, err error) {
	type authorizationFunc func(ctx context.Context, access string, entityId uint64) (userId uint64, err error)

	authFns := map[string]authorizationFunc{
		"drafts":   authorizeProductsDrafts,
		"products": authorizeProductsProducts,
	}
	authFn, ok := authFns[entityTable]
	if !ok {
		return 0, status.Error(xerrors.InvalidArgument, "invalid entity table given")
	}

	return authFn(ctx, access, entityId)
}

func authorizeProductsDrafts(ctx context.Context, access string, entityId uint64) (uint64, error) {
	c, closer, err := pbconf.ConnProductsCreator(ctx)
	if err != nil {
		return 0, err
	}
	defer closer()
	draft, err := c.RetrieveDraft(ctx, &pbproducts.RetrieveDraftRequest{AccessToken: access})
	if err != nil {
		return 0, err
	}
	if draft.GetDraftId() != entityId {
		return 0, status.Errorf(xerrors.PermissionDenied, "your draft is %d, while trying to edit %d", draft.GetDraftId(), entityId)
	}
	return draft.GetDraft().GetUserId(), nil
}

func authorizeProductsProducts(ctx context.Context, access string, entityId uint64) (uint64, error) {
	auth, authCloser, err := pbconf.ConnAuth(ctx)
	if err != nil {
		return 0, err
	}
	userId, err := GetUserFromAccessToken(ctx, auth, access)
	if err != nil {
		return 0, err
	}
	authCloser()
	c, closer, err := pbconf.ConnProductsViewer(ctx)
	if err != nil {
		return 0, err
	}
	defer closer()
	prod, err := c.RetrieveProductDetail(ctx, &pbproducts.RetrieveProductDetailRequest{ProductId: entityId})
	if err != nil {
		return 0, err
	}
	if prod.GetProduct().GetUserId() != userId {
		return 0, status.Error(xerrors.PermissionDenied, "product isnt yours")
	}
	return userId, nil
}
