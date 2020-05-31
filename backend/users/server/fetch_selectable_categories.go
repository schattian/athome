package server

import (
	"context"
	"database/sql"

	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/athome/backend/users/pb/pbuser"
	"github.com/athomecomar/semantic/semprov"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

func (s *Server) FetchSelectableCategories(ctx context.Context, in *pbuser.FetchSelectableCategoriesRequest) (*pbuser.FetchSelectableCategoriesResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := connDB()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "connDB: %v", err)
	}
	defer db.Close()
	return s.fetchSelectableCategories(ctx, db, in)
}

func (s *Server) fetchSelectableCategories(ctx context.Context, db *sqlx.DB, in *pbuser.FetchSelectableCategoriesRequest) (out *pbuser.FetchSelectableCategoriesResponse, err error) {
	onboarding, err := fetchOnboardingByToken(ctx, db, in.GetOnboardingId())
	if errors.Is(err, sql.ErrNoRows) {
		err = status.Errorf(xerrors.NotFound, "onboarding with id %v not found", in.GetOnboardingId())
		return
	}
	if err != nil {
		err = status.Errorf(xerrors.Internal, "fetchOnboardingByToken: %v", err)
		return
	}

	switch onboarding.Role {
	case field.Merchant:
		out, err = fetchSelectableCategoriesMerchant()
	case field.ServiceProvider:
		out = fetchSelectableCategoriesServiceProvider()
	default:
		err = status.Errorf(xerrors.InvalidArgument, "invalid role given: %v", onboarding.Role)
	}
	return
}

func fetchSelectableCategoriesServiceProvider() (out *pbuser.FetchSelectableCategoriesResponse) {
	var cats []*pbuser.Category
	for _, cat := range semprov.Root.Childs {
		cats = append(cats, semprovCategoryToPbCategory(cat))
	}
	return &pbuser.FetchSelectableCategoriesResponse{
		Categories: cats,
	}
}

func semprovCategoryToPbCategory(c *semprov.Category) *pbuser.Category {
	var childs []*pbuser.Category
	for _, child := range c.Childs {
		childs = append(childs, semprovCategoryToPbCategory(child))
	}
	return &pbuser.Category{Name: c.Name, Childs: childs}
}

func fetchSelectableCategoriesMerchant() (out *pbuser.FetchSelectableCategoriesResponse, err error) {
	return nil, status.Error(xerrors.Unimplemented, "not implemented yet")
}
