package ent

type ServiceProviderCategory struct {
	Id uint64 `json:"id,omitempty"`

	Name     string `json:"name,omitempty"`
	ParentId uint64 `json:"parent_id,omitempty"`
}
