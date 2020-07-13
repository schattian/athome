package srvpayments

import "github.com/athomecomar/athome/backend/checkout/internal/checkouttest"

var (
	gPurchases    checkouttest.GoldenPurchases
	gShippings    checkouttest.GoldenShippings
	gStateChanges checkouttest.GoldenStateChanges

	gPbProducts  checkouttest.GoldenPbProducts
	gPbUsers     checkouttest.GoldenPbUsers
	gPbAddresses checkouttest.GoldenPbAddresses

	gPayments checkouttest.GoldenPayments
	gCards    checkouttest.GoldenCards
)

func init() {
	checkouttest.Init(&gPurchases, &gShippings, &gPbProducts, &gStateChanges, &gPbUsers, &gPbAddresses, &gCards, &gPayments)
}
