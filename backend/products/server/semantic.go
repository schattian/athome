package server

import (
	"context"

	"github.com/athomecomar/athome/backend/products/pb/pbproducts"
	"github.com/athomecomar/athome/backend/products/pb/pbsemantic"
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

func PbSemanticToPbProductAttributes(att *pbsemantic.AttributeData) *pbproducts.AttributeData {
	return &pbproducts.AttributeData{
		SchemaId: att.GetSchemaId(),
		Values:   att.GetValues(),
	}
}

func PbSemanticGetAttributesDataToPbProductAttributes(r *pbsemantic.GetAttributesDataResponse) (atts []*pbproducts.AttributeData) {
	for _, att := range r.GetAttributes() {
		atts = append(atts, PbSemanticToPbProductAttributes(att.Data))
	}
	return atts
}
