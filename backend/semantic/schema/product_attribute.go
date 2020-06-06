package schema

import (
	"context"

	"github.com/athomecomar/athome/backend/semantic/data"
	"github.com/athomecomar/athome/backend/semantic/data/value"
	"github.com/jmoiron/sqlx"
)

type ProductAttributeSchema struct {
	Id uint64 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	ValueType value.Type `json:"value_type,omitempty"`

	CategoryId uint64 `json:"category_id,omitempty"`
}

func (pc *ProductAttributeSchema) SetValueType(v value.Type) {
	pc.ValueType = v
}

func (pc *ProductAttributeSchema) GetValueType() value.Type {
	return pc.ValueType
}

func (pc *ProductAttributeSchema) GetName() string {
	return pc.Name
}

func (pc *ProductAttributeSchema) NewData() (data.Attribute, error) {
	att, err := data.NewProductAttributeData(pc.ValueType)
	if err != nil {
		return nil, err
	}
	return att, nil
}

func FindProductAttributeSchema(ctx context.Context, db *sqlx.DB, id uint64) (*ProductAttributeSchema, error) {
	row := db.QueryRowxContext(ctx, `SELECT * FROM product_attributes WHERE id=$1`, id)
	att := &ProductAttributeSchema{}
	err := row.StructScan(att)
	if err != nil {
		return nil, err
	}
	return att, nil
}

func (pc *ProductAttributeSchema) SetName(s string) {
	pc.Name = s
}

func (pc *ProductAttributeSchema) RetrieveCategoryId() uint64 {
	return pc.CategoryId
}

func (pc *ProductAttributeSchema) SetCategoryId(p uint64) {
	pc.CategoryId = p
}
