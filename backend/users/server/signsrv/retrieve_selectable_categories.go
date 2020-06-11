package signsrv

import (
	"context"

	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/athome/backend/users/pb/pbusers"
	"github.com/athomecomar/athome/backend/users/server"
	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) RetrieveSelectableCategories(ctx context.Context, in *pbusers.RetrieveSelectableCategoriesRequest) (*pbusers.RetrieveSelectableCategoriesResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "server.ConnDB: %v", err)
	}
	o, err := retrieveLatestOnboarding(ctx, db, in.GetOnboardingId())
	if err != nil {
		return nil, err
	}
	db.Close()

	return s.retrieveSelectableCategories(ctx, o)
}

func (s *Server) retrieveSelectableCategories(
	ctx context.Context,
	onboarding *ent.Onboarding,
) (out *pbusers.RetrieveSelectableCategoriesResponse, err error) {

	sem, semCloser, err := server.ConnCategories(ctx, onboarding.Role)
	if err != nil {
		return
	}
	defer semCloser()

	categories, err := sem.RetrieveCategories(ctx, &emptypb.Empty{})
	if err != nil {
		return
	}

	out = &pbusers.RetrieveSelectableCategoriesResponse{}
	out.Categories = make(map[uint64]*pbusers.Category)
	for id, cat := range categories.Categories {
		out.Categories[id] = server.PbSemanticCategoryToPbUserCategory(cat)
	}
	return
}
