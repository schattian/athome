package server

// func ConnSemantic(ctx context.Context) (pbsemantic.ProductsClient, func() error, error) {
// 	conn, err := grpc.Dial(addressconf.GetSEMANTIC_ADDR(), grpc.WithInsecure(), grpc.WithBlock())

// 	if err != nil {
// 		return nil, nil, status.Errorf(xerrors.Internal, "grpc.Dial: %v at %v", err, addressconf.GetSEMANTIC_ADDR())
// 	}
// 	c := pbsemantic.NewaddressProvidersClient(conn)
// 	return c, conn.Close, nil
// }

// func ConnImages(ctx context.Context) (pbimages.ImagesClient, func() error, error) {
// 	conn, err := grpc.Dial(addressconf.GetIMAGES_ADDR(), grpc.WithInsecure(), grpc.WithBlock())

// 	if err != nil {
// 		return nil, nil, status.Errorf(xerrors.Internal, "grpc.Dial: %v at %v", err, addressconf.GetIMAGES_ADDR())
// 	}
// 	c := pbimages.NewImagesClient(conn)
// 	return c, conn.Close, nil
// }

// func ConnUsers(ctx context.Context) (pbusers.ViewerClient, func() error, error) {
// 	conn, err := grpc.Dial(addressconf.GetUSERS_ADDR(), grpc.WithInsecure(), grpc.WithBlock())

// 	if err != nil {
// 		return nil, nil, status.Errorf(xerrors.Internal, "grpc.Dial: %v at %v", err, addressconf.GetUSERS_ADDR())
// 	}
// 	c := pbusers.NewViewerClient(conn)
// 	return c, conn.Close, nil
// }

// func PbSemanticToPbProductAttributes(att *pbsemantic.AttributeData) *pbaddress.AttributeData {
// 	return &pbaddress.AttributeData{
// 		SchemaId: att.GetSchemaId(),
// 		Values:   att.GetValues(),
// 	}
// }

// func PbSemanticRetrieveAttributesDataToPbProductAttributes(r *pbsemantic.RetrieveAttributesDataResponse) (atts []*pbaddress.AttributeData) {
// 	for _, att := range r.GetAttributes() {
// 		atts = append(atts, PbSemanticToPbProductAttributes(att.Data))
// 	}
// 	return atts
// }
