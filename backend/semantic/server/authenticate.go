package server

import (
	"context"

	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/athome/pb/pbsemantic"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc/status"
)

func AuthorizeThroughEntity(ctx context.Context, access string, entity *pbsemantic.Entity) (userId uint64, err error) {
	type authorizationFunc func(ctx context.Context, access string, entityId uint64) (userId uint64, err error)

	authFns := map[string]authorizationFunc{
		"drafts": authorizeProductsDrafts,
	}
	authFn, ok := authFns[entity.GetEntityTable()]
	if !ok {
		return 0, status.Error(xerrors.InvalidArgument, "invalid entity table given")
	}

	return authFn(ctx, access, entity.GetEntityId())
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
