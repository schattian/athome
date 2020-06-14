package srvviewer

import "github.com/athomecomar/athome/backend/products/internal/productstest"

var (
	gProducts   productstest.GoldenProducts
	gDraftLines productstest.GoldenDraftLines
	gDrafts     productstest.GoldenDrafts
	gPbUsers    productstest.GoldenPbUsers
	gPbImages   productstest.GoldenPbImages
)

func init() {
	productstest.Init(&gDraftLines, &gDrafts, &gProducts, &gPbUsers, &gPbImages)
}
