package ent

func (o *Onboarding) GetId() uint64 {
	return o.Id
}

func (o *Onboarding) SetId(id uint64) {
	o.Id = id
}

func (o *Onboarding) SQLTable() string {
	return "onboardings"
}

func (o *Onboarding) SQLColumns() []string {
	return []string{
		"id",
		"email",
		"stage",
		"role",
		"name",
		"surname",
	}
}