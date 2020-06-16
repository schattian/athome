package messagertest

import (
	"github.com/athomecomar/athome/pb/pbusers"
)

type variadicPbUsers struct {
	Foo *pbusers.User `json:"foo,omitempty"`
	Bar *pbusers.User `json:"bar,omitempty"`
}
