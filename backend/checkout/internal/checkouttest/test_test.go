package checkouttest

var (
	gPurchases    GoldenPurchases
	gShippings    GoldenShippings
	gPbProducts   GoldenPbProducts
	gStateChanges GoldenStateChanges
	gPbUsers      GoldenPbUsers
	gPbAddresses  GoldenPbAddresses
)

func init() {
	Init(&gPurchases, &gShippings, &gPbProducts, &gStateChanges, &gPbUsers, &gPbAddresses)
}
