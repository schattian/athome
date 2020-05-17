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
		if err := tt.email.Validate(); (err != nil) != tt.err {
			t.Fatalf("email.Validate failed: %v got", err)
		}
	}
}
