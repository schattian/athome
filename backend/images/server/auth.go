package server

import (
	"context"

	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc/status"
)

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
	c, closer, err := pbutil.ConnProductsCreator(ctx)
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
	auth, authCloser, err := pbutil.ConnAuth(ctx)
	if err != nil {
		return 0, err
	}
	userId, err := pbutil.GetUserFromAccessToken(ctx, auth, access)
	if err != nil {
		return 0, err
	}
	authCloser()
	c, closer, err := pbutil.ConnProductsViewer(ctx)
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
