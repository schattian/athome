package server

import (
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/athome/pb/pbsemantic"
)

func PbSemanticToPbProductAttributes(att *pbsemantic.AttributeData) *pbproducts.AttributeData {
	return &pbproducts.AttributeData{
		SchemaId: att.GetSchemaId(),
		Values:   att.GetValues(),
	}
}

func PbSemanticRetrieveAttributesDataToPbProductAttributes(r *pbsemantic.RetrieveAttributeDatasResponse) (atts []*pbproducts.AttributeData) {
	for _, data := range r.GetAttributes() {
		atts = append(atts, PbSemanticToPbProductAttributes(data))
	}
	return atts
}
