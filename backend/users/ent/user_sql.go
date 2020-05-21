package ent

func (u *User) GetId() uint64 {
	return u.Id
}

func (u *User) SetId(id uint64) {
	u.Id = id
}

func (u *User) SQLTable() string {
	return "users"
}

func (u *User) SQLColumns() []string {
	return []string{
		"id",
		"email",
		"password_hash",
		"role",
		"name",
		"surname",
	}
}
