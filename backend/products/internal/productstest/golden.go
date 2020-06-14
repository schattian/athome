package productstest

import (
	"github.com/athomecomar/athome/backend/products/ent"
	"github.com/athomecomar/athome/pb/pbimages"
	"github.com/athomecomar/xtest/xload"
)

func Init(gDraftLines *GoldenDraftLines, gDrafts *GoldenDrafts, gProducts *GoldenProducts, gPbUsers *GoldenPbUsers, gPbImages *GoldenPbImages) {
	xload.DecodeJsonnet("drafts", gDrafts)
	xload.DecodeJsonnet("products", gProducts)
	xload.DecodeJsonnet("users", gPbUsers)
	xload.DecodeJsonnet("images", gPbImages)
	xload.DecodeJsonnet("draft_lines", gDraftLines)
}

type GoldenDraftLines struct {
	Foo *variadicDraftLines `json:"foo,omitempty"`
	Bar *variadicDraftLines `json:"bar,omitempty"`
}

type GoldenPbImages struct {
	Foo *pbimages.Image
	Bar *pbimages.Image
}

type GoldenPbUsers struct {
	Consumers *variadicPbUsers `json:"consumers,omitempty"`
	Merchants *variadicPbUsers `json:"merchants,omitempty"`
}
type GoldenDrafts struct {
	Foo *ent.Draft `json:"foo,omitempty"`
	Bar *ent.Draft `json:"bar,omitempty"`
}

type GoldenProducts struct {
	Foo *variadicProducts `json:"foo,omitempty"`
	Bar *variadicProducts `json:"bar,omitempty"`
}
