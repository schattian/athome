package ent

import (
	"context"
	"fmt"

	"github.com/athomecomar/athome/backend/products/pb/pbimages"
	"github.com/athomecomar/athome/backend/products/pb/pbproducts"
	"github.com/athomecomar/athome/backend/products/pb/pbsemantic"
	"github.com/athomecomar/athome/backend/products/pb/pbusers"
	"github.com/athomecomar/currency"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Product struct {
	Id     uint64 `json:"id,omitempty"`
	UserId uint64 `json:"user_id,omitempty"`

	Title      string `json:"title,omitempty"`
	CategoryId uint64 `json:"category_id,omitempty"`

	Price currency.ARS `json:"price,omitempty"`
	Stock uint64       `json:"stock,omitempty"`

	ImageIds []string `json:"image_ids,omitempty"`
}

func FindProduct(ctx context.Context, db *sqlx.DB, id uint64) (*Product, error) {
	row := db.QueryRowxContext(ctx, `SELECT * FROM products WHERE id=$1`, id)
	prod := &Product{}
	err := row.StructScan(prod)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return prod, nil
}

func (p *Product) GetUser(ctx context.Context, users pbusers.ViewerClient) (*pbproducts.User, error) {
	resp, err := users.ViewUser(ctx, &pbusers.ViewUserRequest{UserId: p.UserId})
	if err != nil {
		return nil, errors.Wrap(err, "ViewUser")
	}
	return &pbproducts.User{Name: resp.GetName(), Surname: resp.GetSurname()}, nil
}

func (p *Product) GetImages(ctx context.Context, img pbimages.ImagesClient) ([]string, error) {
	resp, err := img.RetrieveImages(ctx, &pbimages.RetrieveImagesRequest{Ids: p.ImageIds})
	if err != nil {
		return nil, errors.Wrap(err, "RetrieveImages")
	}
	var uris []string
	for _, image := range resp.GetImages() {
		uris = append(uris, image.GetUri())
	}
	return uris, nil
}

func (p *Product) GetViewableAttributes(ctx context.Context, sem pbsemantic.ProductsClient) ([]*pbproducts.ViewableAttributeData, error) {
	schemas, err := sem.RetrieveAttributesSchema(ctx, &pbsemantic.RetrieveAttributesSchemaRequest{CategoryId: p.CategoryId})
	if err != nil {
		return nil, errors.Wrap(err, "sem.RetrieveAttributesSchema")
	}
	datas, err := sem.RetrieveAttributesData(ctx, &pbsemantic.RetrieveAttributesDataRequest{EntityId: p.Id, EntityTable: p.SQLTable()})
	if err != nil {
		return nil, errors.Wrap(err, "sem.RetrieveAttributesData")
	}
	var atts []*pbproducts.ViewableAttributeData
	for _, data := range datas.GetAttributes() {
		schema := schemaById(data.GetData().GetSchemaId(), schemas.GetAttributes())
		if schema == nil {
			return nil, fmt.Errorf("couldnt find schema with id: %v for attribute data: %v", data.GetData().GetSchemaId(), data.GetAttributeDataId())
		}
		atts = append(atts, attributeFromDataAndSchema(data.GetData(), schema))
	}

	return atts, nil
}

func attributeFromDataAndSchema(data *pbsemantic.AttributeData, schema *pbsemantic.AttributeSchema) *pbproducts.ViewableAttributeData {
	return &pbproducts.ViewableAttributeData{
		Name:      schema.GetName(),
		ValueType: schema.GetValueType(),
		Values:    data.GetValues(),
		SchemaId:  schema.GetId(),
	}
}

func schemaById(id uint64, schemas []*pbsemantic.AttributeSchema) *pbsemantic.AttributeSchema {
	for _, schema := range schemas {
		if schema.GetId() == id {
			return schema
		}
	}

	return nil
}
