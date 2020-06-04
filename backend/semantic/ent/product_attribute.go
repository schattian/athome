package ent

import "github.com/athomecomar/athome/backend/semantic/ent/field"

type ProductAttribute struct {
	Id uint64 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	IsMultivalued bool `json:"is_multivalued,omitempty"`

	Values    []interface{}   `json:"values,omitempty"`
	ValueType field.ValueType `json:"value_type,omitempty"`

	CategoryId uint64 `json:"category_id,omitempty"`
}

func (pc *ProductAttribute) GetIsMultivalued() bool {
	return pc.IsMultivalued
}

func (pc *ProductAttribute) GetValues() []interface{} {
	return pc.Values
}

func (pc *ProductAttribute) SetValues(v ...interface{}) {
	pc.Values = v
}

func (pc *ProductAttribute) GetValueType() field.ValueType {
	return pc.ValueType
}

func (pc *ProductAttribute) SetValueType(v field.ValueType) {
	pc.ValueType = v
}

func (pc *ProductAttribute) GetName() string {
	return pc.Name
}

func (pc *ProductAttribute) SetName(s string) {
	pc.Name = s
}

func (pc *ProductAttribute) GetCategoryId() uint64 {
	return pc.CategoryId
}

func (pc *ProductAttribute) SetCategoryId(p uint64) {
	pc.CategoryId = p
}
