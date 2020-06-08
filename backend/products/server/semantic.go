package server

import (
	"context"

	"github.com/athomecomar/athome/backend/products/pb/pbimages"
	"github.com/athomecomar/athome/backend/products/pb/pbproducts"
	"github.com/athomecomar/athome/backend/products/pb/pbsemantic"
	"github.com/athomecomar/athome/backend/products/pb/pbusers"
	"github.com/athomecomar/athome/backend/products/productconf"
	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func ConnSemantic(ctx context.Context) (pbsemantic.ProductsClient, func() error, error) {
	conn, err := grpc.Dial(productconf.GetSEMANTIC_ADDR(), grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		return nil, nil, status.Errorf(xerrors.Internal, "grpc.Dial: %v at %v", err, productconf.GetSEMANTIC_ADDR())
	}
	c := pbsemantic.NewProductsClient(conn)
	return c, conn.Close, nil
}

func ConnImages(ctx context.Context) (pbimages.ImagesClient, func() error, error) {
	conn, err := grpc.Dial(productconf.GetIMAGES_ADDR(), grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		return nil, nil, status.Errorf(xerrors.Internal, "grpc.Dial: %v at %v", err, productconf.GetIMAGES_ADDR())
	}
	c := pbimages.NewImagesClient(conn)
	return c, conn.Close, nil
}

func ConnUsers(ctx context.Context) (pbusers.ViewerClient, func() error, error) {
	conn, err := grpc.Dial(productconf.GetUSERS_ADDR(), grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		return nil, nil, status.Errorf(xerrors.Internal, "grpc.Dial: %v at %v", err, productconf.GetUSERS_ADDR())
	}
	c := pbusers.NewViewerClient(conn)
	return c, conn.Close, nil
}

func PbSemanticToPbProductAttributes(att *pbsemantic.AttributeData) *pbproducts.AttributeData {
	return &pbproducts.AttributeData{
		SchemaId: att.GetSchemaId(),
		Values:   att.GetValues(),
	}
}

func PbSemanticRetrieveAttributesDataToPbProductAttributes(r *pbsemantic.RetrieveAttributesDataResponse) (atts []*pbproducts.AttributeData) {
	for _, att := range r.GetAttributes() {
		atts = append(atts, PbSemanticToPbProductAttributes(att.Data))
	}
	return atts
}
