package srvviewer

import (
	"context"
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/athomecomar/athome/pb/pbimages"
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/athome/pb/pbsemantic"
	"github.com/athomecomar/athome/pb/pbusers"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/athome/pb/test/pbimagestest"
	"github.com/athomecomar/athome/pb/test/pbsemantictest"
	"github.com/athomecomar/athome/pb/test/pbuserstest"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/storeql/test/sqlassist"
	"github.com/athomecomar/storeql/test/sqlhelp"
	"github.com/athomecomar/xerrors"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc/status"
)

func TestServer_retrieveProductDetail(t *testing.T) {
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
	type semDataStub struct {
		req  *pbsemantic.RetrieveAttributeDatasRequest
		resp *pbsemantic.RetrieveAttributeDatasResponse
		err  error
	}
	type semSchemaStub struct {
		req  *pbsemantic.RetrieveAttributeSchemasRequest
		resp *pbsemantic.RetrieveAttributeSchemasResponse
		err  error
	}
	type args struct {
		ctx context.Context
		in  *pbproducts.RetrieveProductDetailRequest
	}
	tests := []struct {
		name string

		users     usersStub
		imgs      imgsStub
		semData   semDataStub
		semSchema semSchemaStub

		queryStubs []*sqlassist.QueryStubber

		args     args
		want     *pbproducts.ProductDetail
		wantCode xerrors.Code
	}{

		{
			name: "basic retrieve detail",
			imgs: imgsStub{
				req: &pbimages.RetrieveImagesRequest{Entity: pbutil.ToPbImagesEntity(gProducts.Foo.A)},
				resp: &pbimages.RetrieveImagesResponse{
					Images: map[string]*pbimages.Image{
						"fooImageId": gPbImages.Foo,
					},
				},
			},
			semData: semDataStub{
				req: &pbsemantic.RetrieveAttributeDatasRequest{Entity: pbutil.ToPbSemanticEntity(gProducts.Foo.A)},
				resp: &pbsemantic.RetrieveAttributeDatasResponse{
					Attributes: map[uint64]*pbsemantic.AttributeData{
						3424: {SchemaId: 432, Values: []string{"1"}},
						322:  {SchemaId: 384, Values: []string{"hi"}},
					},
				},
			},
			semSchema: semSchemaStub{
				req: &pbsemantic.RetrieveAttributeSchemasRequest{CategoryId: gProducts.Foo.A.CategoryId},
				resp: &pbsemantic.RetrieveAttributeSchemasResponse{
					Attributes: map[uint64]*pbsemantic.AttributeSchema{
						432: {Name: "cantidad de cosas", ValueType: "int64"},
						384: {Name: "comentarioxd", ValueType: "string"},
					},
				},
			},
			users: usersStub{
				req:  &pbusers.RetrieveUserRequest{UserId: gProducts.Foo.A.UserId},
				resp: &pbusers.UserDetail{User: &pbusers.User{Name: gPbUsers.Merchants.Foo.Name, Surname: gPbUsers.Merchants.Foo.Surname}},
			},
			args: args{
				ctx: context.Background(),
				in:  &pbproducts.RetrieveProductDetailRequest{ProductId: gProducts.Foo.A.Id},
			},
			queryStubs: []*sqlassist.QueryStubber{
				{
					Expect: "SELECT * FROM products WHERE id=$1",
					Args:   []driver.Value{gProducts.Foo.A.Id},
					Rows:   sqlmock.NewRows(storeql.SQLColumns(gProducts.Foo.A)).AddRow(storeql.SQLValues(gProducts.Foo.A)...),
				},
			},
			want: &pbproducts.ProductDetail{
				Product: gProducts.Foo.A.ToPb(),
				User:    &pbproducts.User{Name: gPbUsers.Merchants.Foo.Name, Surname: gPbUsers.Merchants.Foo.Surname},
				Images: map[string]*pbproducts.Image{
					"fooImageId": {
						Uri: gPbImages.Foo.Uri,
					},
				},
				Attributes: map[uint64]*pbproducts.Attribute{
					3424: {SchemaId: 432, Values: []string{"1"}, ValueType: "int64", Name: "cantidad de cosas"},
					322:  {SchemaId: 384, Values: []string{"hi"}, ValueType: "string", Name: "comentarioxd"},
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
			users.EXPECT().RetrieveUser(tt.args.ctx, tt.users.req).Return(tt.users.resp, tt.users.err)
			imgs := pbimagestest.NewMockImagesClient(ctrl)
			imgs.EXPECT().RetrieveImages(tt.args.ctx, tt.imgs.req).Return(tt.imgs.resp, tt.imgs.err)
			sem := pbsemantictest.NewMockProductsClient(ctrl)
			sem.EXPECT().RetrieveAttributeDatas(tt.args.ctx, tt.semData.req).Return(tt.semData.resp, tt.semData.err)
			sem.EXPECT().RetrieveAttributeSchemas(tt.args.ctx, tt.semSchema.req).Return(tt.semSchema.resp, tt.semSchema.err)

			db, mock := sqlhelp.MockDB(t)
			for _, stub := range tt.queryStubs {
				stub.Stub(mock)
			}

			got, err := s.retrieveProductDetail(tt.args.ctx, db, users, sem, imgs, tt.args.in)
			if gotCode := status.Code(err); gotCode != tt.wantCode {
				t.Errorf("Server.retrieveProductDetail() error = %v, wantCode %v", err, tt.wantCode)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Server.retrieveProductDetail()  errored mismatch (-want +got): %s", diff)
			}
		})
	}
}
