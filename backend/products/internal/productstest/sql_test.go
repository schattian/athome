package productstest

import (
	"testing"

	"github.com/athomecomar/athome/backend/products/ent"
	"github.com/athomecomar/storeql/test/sqltest"
)

func TestSQL(t *testing.T) {
	sqltest.SQL(t, &ent.Product{}, "Product")
	sqltest.SQL(t, &ent.Draft{}, "Draft")
	sqltest.SQL(t, &ent.DraftLine{}, "DraftLine")
}
