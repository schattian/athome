package productstest

import (
	"testing"

	"github.com/athomecomar/athome/backend/products/ent"
)

func CopyDraft(t *testing.T, c *ent.Draft) *ent.Draft {
	t.Helper()
	if c == nil {
		t.Fatal("cant copy nil draft")
	}
	cp := ent.Draft{}
	cp = *c
	return &cp
}
