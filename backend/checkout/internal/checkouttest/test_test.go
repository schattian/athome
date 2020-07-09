package checkouttest

var (
	gPurchases    GoldenPurchases
	gShippings    GoldenShippings
	gPbProducts   GoldenPbProducts
	gStateChanges GoldenStateChanges
)

func init() {
	Init(&gPurchases, &gShippings, &gPbProducts, &gStateChanges)
}
