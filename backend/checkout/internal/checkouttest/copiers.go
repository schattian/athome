package checkouttest

import (
	"testing"

	"github.com/athomecomar/athome/backend/checkout/ent/order/purchase"
	"github.com/athomecomar/athome/backend/checkout/ent/payment"
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

func CopyPayment(t *testing.T, p *payment.Payment) *payment.Payment {
	t.Helper()
	if p == nil {
		t.Fatal("cant copy nil")
	}
	cp := payment.Payment{}
	cp = *p
	return &cp
}
