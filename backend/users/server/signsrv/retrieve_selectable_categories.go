package signsrv

import (
	"context"
	"database/sql"

	"github.com/athomecomar/athome/backend/users/pb/pbsemantic"
	"github.com/athomecomar/athome/backend/users/pb/pbusers"
	"github.com/athomecomar/athome/backend/users/server"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
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
	defer db.Close()

	return s.retrieveSelectableCategories(ctx, db, in)
}

func (s *Server) retrieveSelectableCategories(
	ctx context.Context, db *sqlx.DB,
	in *pbusers.RetrieveSelectableCategoriesRequest,
) (out *pbusers.RetrieveSelectableCategoriesResponse, err error) {
	onboarding, err := retrieveOnboardingByToken(ctx, db, in.GetOnboardingId())
	if errors.Is(err, sql.ErrNoRows) {
		err = status.Errorf(xerrors.NotFound, "onboarding with id %v not found", in.GetOnboardingId())
		return
	}
	if err != nil {
		err = status.Errorf(xerrors.Internal, "retrieveOnboardingByToken: %v", err)
		return
	}

	sem, semCloser, err := server.ConnSemantic(ctx, onboarding.Role)
	if err != nil {
		return
	}
	defer semCloser()

	categories, err := sem.RetrieveCategories(ctx, &emptypb.Empty{})
	if err != nil {
		return
	}

	for _, cat := range categories.Categories {
		out.Categories = append(out.Categories, pbSemanticCategoryToPbUserCategory(cat))
	}
	return
}

func pbSemanticCategoryToPbUserCategory(c *pbsemantic.Category) *pbusers.Category {
	var childs []*pbusers.Category
	for _, child := range c.GetChilds() {
		childs = append(childs, pbSemanticCategoryToPbUserCategory(child))
	}
	return &pbusers.Category{Name: c.GetName(), Childs: childs, Id: c.GetId(), ParentId: c.GetParentId()}
}
