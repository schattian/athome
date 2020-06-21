package viewersrv

import (
	"context"

	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/athome/backend/users/internal/xpbsemantic"
	"github.com/athomecomar/athome/backend/users/server"
	"github.com/athomecomar/athome/pb/pbimages"
	"github.com/athomecomar/athome/pb/pbusers"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) RetrieveUser(ctx context.Context, in *pbusers.RetrieveUserRequest) (*pbusers.UserDetail, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	user, err := server.FindUser(ctx, db, in.GetUserId())
	if err != nil {
		return nil, err
	}

	sem, semCloser, err := server.ConnCategories(ctx, user.Role)
	if err != nil {
		return nil, err
	}
	defer semCloser()

	img, imgCloser, err := pbutil.ConnImages(ctx)
	if err != nil {
		return nil, err
	}
	defer imgCloser()

	return s.retrieveUser(ctx, db, sem, img, user)
}

func (s *Server) retrieveUser(ctx context.Context, db *sqlx.DB, sem xpbsemantic.CategoriesClient, img pbimages.ImagesClient, user *ent.User) (*pbusers.UserDetail, error) {
	cat, err := user.Category(ctx, sem)
	if err != nil {
		return nil, err
	}
	iden, err := user.Identification(ctx, db)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "Identification: %v", err)
	}
	image, err := user.Image(ctx, img)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "Identification: %v", err)
	}

	return &pbusers.UserDetail{
		User:           user.ToPb(),
		Identification: iden.ToPb(),
		Category:       server.PbSemanticCategoryToPbUserCategory(cat),
		ImageUrl:       image.GetUri(),
	}, nil
}
