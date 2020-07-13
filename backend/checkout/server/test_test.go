package server

import "github.com/athomecomar/athome/backend/checkout/internal/checkouttest"

var (
	gPurchases    checkouttest.GoldenPurchases
	gShippings    checkouttest.GoldenShippings
	gStateChanges checkouttest.GoldenStateChanges

	gPbUsers     checkouttest.GoldenPbUsers
	gPbProducts  checkouttest.GoldenPbProducts
	gPbAddresses checkouttest.GoldenPbAddresses
)

func init() {
	checkouttest.Init(&gPurchases, &gShippings, &gPbProducts, &gStateChanges, &gPbUsers, &gPbAddresses)
}
