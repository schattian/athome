package ent

func (u *User) GetId() int64 {
	return u.Id
}

func (u *User) SetId(id int64) {
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
