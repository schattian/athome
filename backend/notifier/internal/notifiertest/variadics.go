package notifiertest

import (
	"github.com/athomecomar/athome/backend/notifier/ent"
	"github.com/athomecomar/athome/pb/pbusers"
)

type variadicPbUsers struct {
	Foo *pbusers.User `json:"foo,omitempty"`
	Bar *pbusers.User `json:"bar,omitempty"`
}

type variadicNotifications struct {
	Foo *ent.Notification `json:"foo,omitempty"`
	Bar *ent.Notification `json:"bar,omitempty"`
}
