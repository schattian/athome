package field

import "testing"

func Test_Email_Validate(t *testing.T) {
	tests := []struct {
		email Email
		err   bool
		name  string
	}{
		{email: "foo@bar.com", err: false, name: "common email"},
		{email: "foobar.com", err: true, name: "domain"},
		{email: "12312foobar.co321m", err: true, name: "random"},
		{email: "", err: true, name: "nil"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := tt.email.Validate(); (err != nil) != tt.err {
				t.Fatalf("email.Validate failed: %v got", err)
			}
		})
	}
}

func Test_Name_Validate(t *testing.T) {
	tests := []struct {
		namefield Name
		err       bool
		name      string
	}{
		{namefield: "asd@asd.com", err: true, name: "common email"},
		{namefield: "foobar.com", err: false, name: "domain"},
		{namefield: "foobar", err: false, name: "normal name"},
		{namefield: "Foo Bar", err: false, name: "normal name with upper and spaces"},
		{namefield: "Fóo Bar", err: false, name: "normal name with upper and spaces and accents"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := tt.namefield.Validate(); (err != nil) != tt.err {
				t.Fatalf("name.Validate failed. Got: %v", err)
			}
		})
	}
}

func Test_Surname_Validate(t *testing.T) {
	surnames := []Surname{
		"foobar.com",
		"foobar",
		"Foo Bar",
		"Fóo Bar",
	}
	for _, sn := range surnames {
		sn := sn
		t.Run(string(sn), func(t *testing.T) {
			sn := sn
			t.Parallel()
			if sn.Validate() != Name(sn).Validate() {
				t.Fatalf("surname.Validate failed. Differs from name.Validate")
			}
		})
	}
}
