package checkouttest

import (
	"testing"

	"github.com/athomecomar/storeql/test/sqltest"
)

func TestSQL(t *testing.T) {
	sqltest.SQL(t, gPurchases.Foo, "Purchase")
	sqltest.SQL(t, gShippings.Foo, "Shipping")
}
