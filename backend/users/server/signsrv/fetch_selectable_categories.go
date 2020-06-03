package signsrv

import (
	"context"
	"database/sql"

	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/athome/backend/users/pb/pbusers"
	"github.com/athomecomar/athome/backend/users/server"
	"github.com/athomecomar/athome/backend/users/userconf"
	"github.com/athomecomar/semantic/semprov"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
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

	conn, err := grpc.Dial(userconf.GetSEMANTIC_ADDR(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "grpc.Dial: %v at %v", err, userconf.GetSEMANTIC_ADDR())
	}
	defer conn.Close()

	return s.fetchSelectableCategories(ctx, db, conn, in)
}

func (s *Server) fetchSelectableCategories(ctx context.Context, db *sqlx.DB, conn *grpc.ClientConn, in *pbusers.FetchSelectableCategoriesRequest) (out *pbusers.FetchSelectableCategoriesResponse, err error) {
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

func fetchSelectableCategoriesServiceProvider() (out *pbusers.FetchSelectableCategoriesResponse) {
	var cats []*pbusers.Category
	for _, cat := range semprov.Root.Childs {
		cats = append(cats, semprovCategoryToPbCategory(cat))
	}
	return &pbusers.FetchSelectableCategoriesResponse{
		Categories: cats,
	}
}

func semprovCategoryToPbCategory(c *semprov.Category) *pbusers.Category {
	var childs []*pbusers.Category
	for _, child := range c.Childs {
		childs = append(childs, semprovCategoryToPbCategory(child))
	}
	return &pbusers.Category{Name: c.Name, Childs: childs}
}

func fetchSelectableCategoriesMerchant() (out *pbusers.FetchSelectableCategoriesResponse, err error) {
	return nil, status.Error(xerrors.Unimplemented, "not implemented yet")
}
