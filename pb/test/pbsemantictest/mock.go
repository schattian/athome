package pbsemantictest

import (
	"context"

	"github.com/athomecomar/pb/pbsemantic"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Client struct {
	RetrieveCategoriesResponse struct {
		Resp *pbsemantic.RetrieveCategoriesResponse
		Err  error
	}
	RetrieveCategoryResponse struct {
		Resp *pbsemantic.Category
		Err  error
	}
}

func (c Client) RetrieveCategory(ctx context.Context, in *pbsemantic.RetrieveCategoryRequest, opts ...grpc.CallOption) (*pbsemantic.Category, error) {
	return c.RetrieveCategoryResponse.Resp, c.RetrieveCategoryResponse.Err
}

func (c Client) RetrieveCategories(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*pbsemantic.RetrieveCategoriesResponse, error) {
	return c.RetrieveCategoriesResponse.Resp, c.RetrieveCategoriesResponse.Err
}
