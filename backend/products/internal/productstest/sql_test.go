package productstest

import (
	"testing"

	"github.com/athomecomar/athome/backend/products/ent"
	"github.com/athomecomar/storeql/test/sqltest"
)

func TestOnboardingIdentificationsSQL(t *testing.T) {
	sqltest.SQL(t, &ent.Draft{}, "Draft")
	// sqltest.SQL(t, &ent.ProductAttribute{}, "ProductCategory")
}
