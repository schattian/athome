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

func (s *Server) FetchSelectableCategories(ctx context.Context, in *pbusers.FetchSelectableCategoriesRequest) (*pbusers.FetchSelectableCategoriesResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "server.ConnDB: %v", err)
	}
	defer db.Close()

	return s.fetchSelectableCategories(ctx, db, in)
}

func (s *Server) fetchSelectableCategories(
	ctx context.Context, db *sqlx.DB,
	in *pbusers.FetchSelectableCategoriesRequest,
) (out *pbusers.FetchSelectableCategoriesResponse, err error) {
	onboarding, err := fetchOnboardingByToken(ctx, db, in.GetOnboardingId())
	if errors.Is(err, sql.ErrNoRows) {
		err = status.Errorf(xerrors.NotFound, "onboarding with id %v not found", in.GetOnboardingId())
		return
	}
	if err != nil {
		err = status.Errorf(xerrors.Internal, "fetchOnboardingByToken: %v", err)
		return
	}

	sem, semCloser, err := server.ConnSemantic(ctx, onboarding.Role)
	if err != nil {
		return
	}
	defer semCloser()

	categories, err := sem.GetCategories(ctx, &emptypb.Empty{})
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
