package productstest

import (
	"github.com/athomecomar/athome/backend/products/ent"
	"github.com/athomecomar/athome/pb/pbusers"
)

type variadicPbUsers struct {
	Foo *pbusers.User `json:"foo,omitempty"`
	Bar *pbusers.User `json:"bar,omitempty"`
}

type variadicProducts struct {
	A *ent.Product
	B *ent.Product
}

type variadicDraftLines struct {
	A *ent.DraftLine
	B *ent.DraftLine
}
