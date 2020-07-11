package checkouttest

var (
	gPurchases    GoldenPurchases
	gShippings    GoldenShippings
	gPbProducts   GoldenPbProducts
	gStateChanges GoldenStateChanges
	gPbUsers      GoldenPbUsers
)

func init() {
	Init(&gPurchases, &gShippings, &gPbProducts, &gStateChanges, &gPbUsers)
}
