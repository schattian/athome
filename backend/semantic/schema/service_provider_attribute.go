package schema

import (
	"context"

	"github.com/athomecomar/athome/backend/semantic/data"
	"github.com/athomecomar/athome/backend/semantic/data/value"
	"github.com/jmoiron/sqlx"
)

type ServiceProviderAttributeSchema struct {
	Id uint64 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	ValueType value.Type `json:"value_type,omitempty"`

	CategoryId uint64 `json:"category_id,omitempty"`
}

func (pc *ServiceProviderAttributeSchema) SetValueType(v value.Type) {
	pc.ValueType = v
}

func (pc *ServiceProviderAttributeSchema) GetValueType() value.Type {
	return pc.ValueType
}

func (pc *ServiceProviderAttributeSchema) GetName() string {
	return pc.Name
}

func (pc *ServiceProviderAttributeSchema) NewData() (data.Attribute, error) {
	att, err := data.NewServiceProviderAttributeData(pc.ValueType)
	if err != nil {
		return nil, err
	}
	return att, nil
}

func FindServiceProviderAttributeSchema(ctx context.Context, db *sqlx.DB, id uint64) (*ServiceProviderAttributeSchema, error) {
	row := db.QueryRowxContext(ctx, `SELECT * FROM service_provider_attributes WHERE id=$1`, id)
	att := &ServiceProviderAttributeSchema{}
	err := row.StructScan(att)
	if err != nil {
		return nil, err
	}
	return att, nil
}

func (pc *ServiceProviderAttributeSchema) SetName(s string) {
	pc.Name = s
}

func (pc *ServiceProviderAttributeSchema) RetrieveCategoryId() uint64 {
	return pc.CategoryId
}

func (pc *ServiceProviderAttributeSchema) SetCategoryId(p uint64) {
	pc.CategoryId = p
}
