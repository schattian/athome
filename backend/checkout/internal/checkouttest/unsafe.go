package checkouttest

import (
	"testing"

	"github.com/athomecomar/athome/backend/checkout/ent/order/purchase"
	"github.com/athomecomar/athome/backend/checkout/ent/payment"
	"github.com/athomecomar/athome/pb/pbcheckout"
)

func PurchaseToPb(t *testing.T, p *purchase.Purchase, amount float64) *pbcheckout.Purchase {
	t.Helper()
	pb, err := p.ToPb(amount)
	if err != nil {
		t.Fatalf("PurchaseToPb: %v", err)
	}
	return pb
}

func PaymentToPb(t *testing.T, p *payment.Payment) *pbcheckout.Payment {
	t.Helper()
	pb, err := p.ToPb()
	if err != nil {
		t.Fatalf("PurchaseToPb: %v", err)
	}
	return pb
}
