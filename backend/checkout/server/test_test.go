package server

import "github.com/athomecomar/athome/backend/checkout/internal/checkouttest"

var (
	gPurchases    checkouttest.GoldenPurchases
	gShippings    checkouttest.GoldenShippings
	gStateChanges checkouttest.GoldenStateChanges

	gPbUsers     checkouttest.GoldenPbUsers
	gPbProducts  checkouttest.GoldenPbProducts
	gPbAddresses checkouttest.GoldenPbAddresses

	gPayments   checkouttest.GoldenPayments
	gCards      checkouttest.GoldenCards
	gPbEvents   checkouttest.GoldenPbEvents
	gPbServices checkouttest.GoldenPbServices
)

func init() {
	checkouttest.Init(&gPurchases, &gShippings, &gPbProducts, &gStateChanges, &gPbUsers, &gPbAddresses, &gCards, &gPayments, &gPbEvents, &gPbServices)
}
