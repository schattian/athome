package productstest

import (
	"testing"

	"github.com/athomecomar/athome/backend/products/ent"
	"github.com/athomecomar/athome/backend/products/ent/stage"
)

func SetStage(t *testing.T, s *ent.Draft, st stage.Stage) *ent.Draft {
	d := CopyDraft(t, s)
	d.Stage = st
	return d
}
