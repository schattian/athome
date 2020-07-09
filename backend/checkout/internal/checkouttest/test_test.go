package checkouttest

var (
	gPurchases  GoldenPurchases
	gShippings  GoldenShippings
	gPbProducts GoldenPbProducts
)

func init() {
	Init(&gPurchases, &gShippings, &gPbProducts)
}
