package server

// func ConnSemantic(ctx context.Context) (pbsemantic.ProductsClient, func() error, error) {
// 	conn, err := grpc.Dial(serviceconf.GetSEMANTIC_ADDR(), grpc.WithInsecure(), grpc.WithBlock())

// 	if err != nil {
// 		return nil, nil, status.Errorf(xerrors.Internal, "grpc.Dial: %v at %v", err, serviceconf.GetSEMANTIC_ADDR())
// 	}
// 	c := pbsemantic.NewServiceProvidersClient(conn)
// 	return c, conn.Close, nil
// }

// func PbSemanticToPbProductAttributes(att *pbsemantic.AttributeData) *pbservices.AttributeData {
// 	return &pbservices.AttributeData{
// 		SchemaId: att.GetSchemaId(),
// 		Values:   att.GetValues(),
// 	}
// }

// func PbSemanticRetrieveAttributesDataToPbProductAttributes(r *pbsemantic.RetrieveAttributesDataResponse) (atts []*pbservices.AttributeData) {
// 	for _, att := range r.GetAttributes() {
// 		atts = append(atts, PbSemanticToPbProductAttributes(att.Data))
// 	}
// 	return atts
// }
