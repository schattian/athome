package servicestest

import (
	"testing"

	"github.com/athomecomar/athome/backend/services/ent"
)

func CopyService(t *testing.T, c *ent.Service) *ent.Service {
	t.Helper()
	if c == nil {
		t.Fatal("cant copy nil svc")
	}
	cp := ent.Service{}
	cp = *c
	return &cp
}
