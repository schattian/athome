package checkouttest

import (
	"testing"

	"github.com/athomecomar/athome/backend/checkout/ent/order/purchase"
)

func PurchaseZeroShippingId(t *testing.T, p *purchase.Purchase) *purchase.Purchase {
	t.Helper()
	cp := CopyPurchase(t, p)
	cp.ShippingId = 0
	return cp
}

func PurchaseZeroDestAddressId(t *testing.T, p *purchase.Purchase) *purchase.Purchase {
	t.Helper()
	cp := CopyPurchase(t, p)
	cp.DestAddressId = 0
	return cp
}

func PurchaseCreation(t *testing.T, p *purchase.Purchase) *purchase.Purchase {
	t.Helper()
	return PurchaseZeroDestAddressId(t, PurchaseZeroShippingId(t, p))
}
