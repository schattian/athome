package usertest

import (
	"testing"

	"github.com/athomecomar/athome/backend/users/ent"
)

func CopyOnboarding(t *testing.T, o *ent.Onboarding) *ent.Onboarding {
	t.Helper()
	if o == nil {
		t.Fatal("cant copy nil onboarding")
	}
	cp := ent.Onboarding{}
	cp = *o
	return &cp
}
