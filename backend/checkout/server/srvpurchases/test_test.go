package srvpurchases

import "github.com/athomecomar/athome/backend/checkout/internal/checkouttest"

var (
	gPurchases checkouttest.GoldenPurchases
)

func init() {
	checkouttest.Init(&gPurchases)
}
