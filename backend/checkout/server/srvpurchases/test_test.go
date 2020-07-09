package srvpurchases

import "github.com/athomecomar/athome/backend/checkout/internal/checkouttest"

var (
	gPurchases    checkouttest.GoldenPurchases
	gShippings    checkouttest.GoldenShippings
	gPbProducts   checkouttest.GoldenPbProducts
	gStateChanges checkouttest.GoldenStateChanges
)

func init() {
	checkouttest.Init(&gPurchases, &gShippings, &gPbProducts, &gStateChanges)
}
