package messagertest

import (
	"github.com/athomecomar/athome/backend/messager/ent"
	"github.com/athomecomar/xtest/xload"
)

func Init(gPbUsers *GoldenPbUsers, gMessages *GoldenMessages) {
	xload.DecodeJsonnet("users", gPbUsers)
	xload.DecodeJsonnet("messages", gMessages)
}

type GoldenPbUsers struct {
	Consumers *variadicPbUsers `json:"consumers,omitempty"`
	Merchants *variadicPbUsers `json:"merchants,omitempty"`
}

type GoldenMessages struct {
	Foo *ent.Message
	Bar *ent.Message
}
