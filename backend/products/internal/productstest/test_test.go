package productstest

var (
	gProducts   GoldenProducts
	gDraftLines GoldenDraftLines
	gDrafts     GoldenDrafts
	gPbUsers    GoldenPbUsers
	gPbImages   GoldenPbImages
)

func init() {
	Init(&gDraftLines, &gDrafts, &gProducts, &gPbUsers, &gPbImages)
}
