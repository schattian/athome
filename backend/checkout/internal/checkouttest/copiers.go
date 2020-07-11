package checkouttest

import (
	"testing"

	"github.com/athomecomar/athome/backend/checkout/ent/order/purchase"
)

func CopyPurchase(t *testing.T, p *purchase.Purchase) *purchase.Purchase {
	t.Helper()
	if p == nil {
		t.Fatal("cant copy nil")
	}
	cp := purchase.Purchase{}
	cp = *p
	return &cp
}
