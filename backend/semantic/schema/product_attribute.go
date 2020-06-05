package schema

import "github.com/athomecomar/athome/backend/semantic/data/value"

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

func (pc *ProductAttributeSchema) SetName(s string) {
	pc.Name = s
}

func (pc *ProductAttributeSchema) GetCategoryId() uint64 {
	return pc.CategoryId
}

func (pc *ProductAttributeSchema) SetCategoryId(p uint64) {
	pc.CategoryId = p
}
