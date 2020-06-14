package srvviewer

import (
	"context"
	"database/sql/driver"
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/athomecomar/athome/pb/pbimages"
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/athome/pb/pbusers"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/athome/pb/test/pbimagestest"
	"github.com/athomecomar/athome/pb/test/pbuserstest"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/storeql/test/sqlassist"
	"github.com/athomecomar/storeql/test/sqlhelp"
	"github.com/athomecomar/xerrors"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc/status"
)

func TestServer_searchProducts(t *testing.T) {
	type usersStub struct {
		req  *pbusers.RetrieveUserRequest
		resp *pbusers.UserDetail
		err  error
	}
	type imgsStub struct {
		req  *pbimages.RetrieveImagesRequest
		resp *pbimages.RetrieveImagesResponse
		err  error
	}
	type args struct {
		ctx context.Context
		in  *pbproducts.SearchProductsRequest
	}
	tests := []struct {
		name       string
		queryStubs []*sqlassist.QueryStubber
		args       args

		users []usersStub
		imgs  []imgsStub

		want     *pbproducts.SearchProductsResponse
		wantCode xerrors.Code
	}{
		{
			name: "nil cursor given (1st iteration)",
			imgs: []imgsStub{
				{
					req: &pbimages.RetrieveImagesRequest{Entity: pbutil.ToPbImagesEntity(gProducts.Foo.A)},
					resp: &pbimages.RetrieveImagesResponse{
						Images: map[string]*pbimages.Image{
							"fooImageId": gPbImages.Foo,
						},
					},
				},
				{
					req: &pbimages.RetrieveImagesRequest{Entity: pbutil.ToPbImagesEntity(gProducts.Bar.A)},
					resp: &pbimages.RetrieveImagesResponse{
						Images: map[string]*pbimages.Image{
							"barImageId": gPbImages.Bar,
						},
					},
				},
				{
					req: &pbimages.RetrieveImagesRequest{Entity: pbutil.ToPbImagesEntity(gProducts.Foo.B)},
					resp: &pbimages.RetrieveImagesResponse{
						Images: map[string]*pbimages.Image{
							"fooImageId": gPbImages.Foo,
						},
					},
				},
			},
			users: []usersStub{
				{
					req:  &pbusers.RetrieveUserRequest{UserId: gProducts.Foo.A.UserId},
					resp: &pbusers.UserDetail{User: gPbUsers.Merchants.Foo},
				},
				{
					req:  &pbusers.RetrieveUserRequest{UserId: gProducts.Bar.A.UserId},
					resp: &pbusers.UserDetail{User: gPbUsers.Merchants.Bar},
				},
				{
					req:  &pbusers.RetrieveUserRequest{UserId: gProducts.Foo.B.UserId},
					resp: &pbusers.UserDetail{User: gPbUsers.Merchants.Foo},
				},
			},
			args: args{
				ctx: context.Background(),
				in: &pbproducts.SearchProductsRequest{
					Query: "FÓóBáR",
					Page: &pbproducts.PageRequest{
						Size: 3,
					},
				},
			},
			queryStubs: []*sqlassist.QueryStubber{
				{
					Expect: `SELECT (*) FROM products 
    WHERE lower(unaccent(title)) ILIKE ESCAPE $1
    ORDER BY id DESC LIMIT $2`,
					Args: []driver.Value{"foobar", 3},
					Rows: sqlmock.NewRows(storeql.SQLColumns(gProducts.Foo.A)).
						AddRow(storeql.SQLValues(gProducts.Foo.A)...).
						AddRow(storeql.SQLValues(gProducts.Bar.A)...).
						AddRow(storeql.SQLValues(gProducts.Foo.B)...),
				},
				{
					Expect: "SELECT COUNT(*) FROM products WHERE lower(unaccent(title)) ILIKE ESCAPE $1",
					Args:   []driver.Value{"foobar"},
					Rows:   sqlmock.NewRows([]string{"COUNT(*)"}).AddRow(4),
				},
			},
			want: &pbproducts.SearchProductsResponse{
				Page: &pbproducts.PageResponse{
					NextCursor: b64EncodeId(gProducts.Foo.B.Id),
					TotalSize:  4,
				},
				Products: map[uint64]*pbproducts.ProductSearchResult{
					gProducts.Foo.A.Id: {
						Product: &pbproducts.ProductSearchResult_Product{
							Title: gProducts.Foo.A.Title,
							Price: gProducts.Foo.A.Price.Float64(),
						},
						User: &pbproducts.User{Name: gPbUsers.Merchants.Foo.Name, Surname: gPbUsers.Merchants.Foo.Surname},
						Images: map[string]*pbproducts.Image{
							"fooImageId": {Uri: gPbImages.Foo.Uri},
						},
					},

					gProducts.Bar.A.Id: {
						Product: &pbproducts.ProductSearchResult_Product{
							Title: gProducts.Bar.A.Title,
							Price: gProducts.Bar.A.Price.Float64(),
						},
						User: &pbproducts.User{Name: gPbUsers.Merchants.Bar.Name, Surname: gPbUsers.Merchants.Bar.Surname},
						Images: map[string]*pbproducts.Image{
							"barImageId": {Uri: gPbImages.Bar.Uri},
						},
					},

					gProducts.Foo.B.Id: {
						Product: &pbproducts.ProductSearchResult_Product{
							Title: gProducts.Foo.B.Title,
							Price: gProducts.Foo.B.Price.Float64(),
						},
						User: &pbproducts.User{Name: gPbUsers.Merchants.Foo.Name, Surname: gPbUsers.Merchants.Foo.Surname},
						Images: map[string]*pbproducts.Image{
							"fooImageId": {Uri: gPbImages.Foo.Uri},
						},
					},
				},
			},
		},

		{
			name: "basic retrieve detail",
			imgs: []imgsStub{
				{
					req: &pbimages.RetrieveImagesRequest{Entity: pbutil.ToPbImagesEntity(gProducts.Foo.A)},
					resp: &pbimages.RetrieveImagesResponse{
						Images: map[string]*pbimages.Image{
							"fooImageId": gPbImages.Foo,
						},
					},
				},
				{
					req: &pbimages.RetrieveImagesRequest{Entity: pbutil.ToPbImagesEntity(gProducts.Bar.A)},
					resp: &pbimages.RetrieveImagesResponse{
						Images: map[string]*pbimages.Image{
							"barImageId": gPbImages.Bar,
						},
					},
				},
				{
					req: &pbimages.RetrieveImagesRequest{Entity: pbutil.ToPbImagesEntity(gProducts.Foo.B)},
					resp: &pbimages.RetrieveImagesResponse{
						Images: map[string]*pbimages.Image{
							"fooImageId": gPbImages.Foo,
						},
					},
				},
			},
			users: []usersStub{
				{
					req:  &pbusers.RetrieveUserRequest{UserId: gProducts.Foo.A.UserId},
					resp: &pbusers.UserDetail{User: gPbUsers.Merchants.Foo},
				},
				{
					req:  &pbusers.RetrieveUserRequest{UserId: gProducts.Bar.A.UserId},
					resp: &pbusers.UserDetail{User: gPbUsers.Merchants.Bar},
				},
				{
					req:  &pbusers.RetrieveUserRequest{UserId: gProducts.Foo.B.UserId},
					resp: &pbusers.UserDetail{User: gPbUsers.Merchants.Foo},
				},
			},
			args: args{
				ctx: context.Background(),
				in: &pbproducts.SearchProductsRequest{
					Query: "FÓóBáR",
					Page: &pbproducts.PageRequest{
						Cursor: b64EncodeId(gProducts.Foo.A.Id + 1),
						Size:   3,
					},
				},
			},
			queryStubs: []*sqlassist.QueryStubber{
				{
					Expect: `SELECT (*) FROM products 
    WHERE lower(unaccent(title)) ILIKE ESCAPE $1
    AND id < ` + strconv.Itoa(int(gProducts.Foo.A.Id)+1) +
						` ORDER BY id DESC LIMIT $2`,
					Args: []driver.Value{"foobar", 3},
					Rows: sqlmock.NewRows(storeql.SQLColumns(gProducts.Foo.A)).
						AddRow(storeql.SQLValues(gProducts.Foo.A)...).
						AddRow(storeql.SQLValues(gProducts.Bar.A)...).
						AddRow(storeql.SQLValues(gProducts.Foo.B)...),
				},
				{
					Expect: "SELECT COUNT(*) FROM products WHERE lower(unaccent(title)) ILIKE ESCAPE $1",
					Args:   []driver.Value{"foobar"},
					Rows:   sqlmock.NewRows([]string{"COUNT(*)"}).AddRow(4),
				},
			},
			want: &pbproducts.SearchProductsResponse{
				Page: &pbproducts.PageResponse{
					NextCursor: b64EncodeId(gProducts.Foo.B.Id),
					TotalSize:  4,
				},
				Products: map[uint64]*pbproducts.ProductSearchResult{
					gProducts.Foo.A.Id: {
						Product: &pbproducts.ProductSearchResult_Product{
							Title: gProducts.Foo.A.Title,
							Price: gProducts.Foo.A.Price.Float64(),
						},
						User: &pbproducts.User{Name: gPbUsers.Merchants.Foo.Name, Surname: gPbUsers.Merchants.Foo.Surname},
						Images: map[string]*pbproducts.Image{
							"fooImageId": {Uri: gPbImages.Foo.Uri},
						},
					},

					gProducts.Bar.A.Id: {
						Product: &pbproducts.ProductSearchResult_Product{
							Title: gProducts.Bar.A.Title,
							Price: gProducts.Bar.A.Price.Float64(),
						},
						User: &pbproducts.User{Name: gPbUsers.Merchants.Bar.Name, Surname: gPbUsers.Merchants.Bar.Surname},
						Images: map[string]*pbproducts.Image{
							"barImageId": {Uri: gPbImages.Bar.Uri},
						},
					},

					gProducts.Foo.B.Id: {
						Product: &pbproducts.ProductSearchResult_Product{
							Title: gProducts.Foo.B.Title,
							Price: gProducts.Foo.B.Price.Float64(),
						},
						User: &pbproducts.User{Name: gPbUsers.Merchants.Foo.Name, Surname: gPbUsers.Merchants.Foo.Surname},
						Images: map[string]*pbproducts.Image{
							"fooImageId": {Uri: gPbImages.Foo.Uri},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := &Server{}
			ctrl := gomock.NewController(t)
			users := pbuserstest.NewMockViewerClient(ctrl)
			for _, user := range tt.users {
				users.EXPECT().RetrieveUser(tt.args.ctx, user.req).Return(user.resp, user.err)
			}
			imgs := pbimagestest.NewMockImagesClient(ctrl)
			for _, img := range tt.imgs {
				imgs.EXPECT().RetrieveImages(tt.args.ctx, img.req).Return(img.resp, img.err)
			}
			db, mock := sqlhelp.MockDB(t)
			for _, stub := range tt.queryStubs {
				stub.Stub(mock)
			}
			got, err := s.searchProducts(tt.args.ctx, db, nil, users, imgs, tt.args.in)
			if gotCode := status.Code(err); gotCode != tt.wantCode {
				t.Errorf("Server.searchProducts() error = %v, wantCode %v", err, tt.wantCode)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Server.searchProducts()  errored mismatch (-want +got): %s", diff)
			}

		})
	}
}
