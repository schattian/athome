package xpbsemantic

import (
	"context"

	"github.com/athomecomar/athome/pb/pbsemantic"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CategoriesClient interface {
	RetrieveCategory(ctx context.Context, in *pbsemantic.RetrieveCategoryRequest, opts ...grpc.CallOption) (*pbsemantic.Category, error)
	RetrieveCategories(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*pbsemantic.RetrieveCategoriesResponse, error)
}
