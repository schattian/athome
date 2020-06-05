package schema

type MerchantCategory struct {
	Id uint64 `json:"id,omitempty"`

	Name     string `json:"name,omitempty"`
	ParentId uint64 `json:"parent_id,omitempty"`
}

func (pc *MerchantCategory) GetName() string {
	return pc.Name
}

func (pc *MerchantCategory) SetName(s string) {
	pc.Name = s
}

func (pc *MerchantCategory) GetParentId() uint64 {
	return pc.ParentId
}

func (pc *MerchantCategory) SetParentId(p uint64) {
	pc.ParentId = p
}
