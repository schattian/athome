package ent

import (
	"context"
	"fmt"

	"github.com/athomecomar/athome/pb/pbimages"
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/athome/pb/pbsemantic"
	"github.com/athomecomar/athome/pb/pbusers"
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

func (p *Product) ToPb() *pbproducts.Product {
	return &pbproducts.Product{
		Title:      p.Title,
		CategoryId: p.CategoryId,
		Price:      p.Price.Float64(),
		Stock:      p.Stock,
		ImageIds:   p.ImageIds,
	}
}

func (p *Product) ToPbSearchResult(ctx context.Context, users pbusers.ViewerClient, img pbimages.ImagesClient) (*pbproducts.ProductSearchResult, error) {
	var err error
	resp := &pbproducts.ProductSearchResult{Product: &pbproducts.ProductSearchResult_Product{Title: p.Title, Price: p.Price.Float64()}}
	resp.User, err = p.GetUser(ctx, users)
	if err != nil {
		return nil, errors.Wrap(err, "GetUser")
	}
	resp.Images, err = p.GetImages(ctx, img)
	if err != nil {
		return nil, errors.Wrap(err, "GetImages")
	}
	return resp, nil
}

func (p *Product) GetUser(ctx context.Context, users pbusers.ViewerClient) (*pbproducts.User, error) {
	resp, err := users.RetrieveUser(ctx, &pbusers.RetrieveUserRequest{UserId: p.UserId})
	if err != nil {
		return nil, errors.Wrap(err, "ViewUser")
	}
	return &pbproducts.User{Name: resp.GetUser().GetName(), Surname: resp.GetUser().GetSurname()}, nil
}

func (p *Product) GetImages(ctx context.Context, img pbimages.ImagesClient) (map[string]*pbproducts.Image, error) {
	resp, err := img.RetrieveImages(ctx, &pbimages.RetrieveImagesRequest{Ids: p.ImageIds})
	if err != nil {
		return nil, errors.Wrap(err, "RetrieveImages")
	}

	images := make(map[string]*pbproducts.Image)
	for id, image := range resp.GetImages() {
		images[id] = &pbproducts.Image{Uri: image.Uri}
	}
	return images, nil
}

func (p *Product) GetAttributes(ctx context.Context, sem pbsemantic.ProductsClient) (map[uint64]*pbproducts.Attribute, error) {
	schemas, err := sem.RetrieveAttributeSchemas(ctx, &pbsemantic.RetrieveAttributeSchemasRequest{CategoryId: p.CategoryId})
	if err != nil {
		return nil, errors.Wrap(err, "sem.RetrieveAttributesSchema")
	}
	datas, err := sem.RetrieveAttributeDatas(ctx, &pbsemantic.RetrieveAttributeDatasRequest{EntityId: p.Id, EntityTable: p.SQLTable()})
	if err != nil {
		return nil, errors.Wrap(err, "sem.RetrieveAttributesData")
	}
	atts := make(map[uint64]*pbproducts.Attribute)
	for id, data := range datas.GetAttributes() {
		schema, ok := schemas.GetAttributes()[data.GetSchemaId()]
		if !ok {
			return nil, fmt.Errorf("couldnt find schema with id: %v for attribute data: %v", data.GetSchemaId(), id)
		}
		atts[id] = attributeFromDataAndSchema(data, schema)
	}

	return atts, nil
}

func attributeFromDataAndSchema(data *pbsemantic.AttributeData, schema *pbsemantic.AttributeSchema) *pbproducts.Attribute {
	return &pbproducts.Attribute{
		Name:      schema.GetName(),
		ValueType: schema.GetValueType(),
		Values:    data.GetValues(),
	}
}
