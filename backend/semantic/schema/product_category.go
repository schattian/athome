package schema

type ProductCategory struct {
	Id uint64 `json:"id,omitempty"`

	Name     string `json:"name,omitempty"`
	ParentId uint64 `json:"parent_id,omitempty"`
}

func (pc *ProductCategory) GetName() string {
	return pc.Name
}

func (pc *ProductCategory) SetName(s string) {
	pc.Name = s
}

func (pc *ProductCategory) GetParentId() uint64 {
	return pc.ParentId
}

func (pc *ProductCategory) SetParentId(p uint64) {
	pc.ParentId = p
}
