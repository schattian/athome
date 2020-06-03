package ent

type ServiceProviderCategory struct {
	Id uint64 `json:"id,omitempty"`

	Name     string `json:"name,omitempty"`
	ParentId uint64 `json:"parent_id,omitempty"`
}

func (pc *ServiceProviderCategory) GetName() string {
	return pc.Name
}

func (pc *ServiceProviderCategory) SetName(s string) {
	pc.Name = s
}

func (pc *ServiceProviderCategory) GetParentId() uint64 {
	return pc.ParentId
}

func (pc *ServiceProviderCategory) SetParentId(p uint64) {
	pc.ParentId = p
}
