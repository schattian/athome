package productstest

import (
	"testing"

	"github.com/athomecomar/storeql/test/sqltest"
)

func TestSQL(t *testing.T) {
	sqltest.SQL(t, gProducts.Foo.A, "Product")
	sqltest.SQL(t, gDrafts.Foo, "Draft")
	sqltest.SQL(t, gDraftLines.Foo.A, "DraftLine")
}
