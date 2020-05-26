package ent

import (
	"github.com/athomecomar/xtest/xload"
)

var (
	gUsers       *goldenUsers
	gOnboardings *goldenOnboardings
)

func init() {
	xload.DecodeJsonnet("users", &gUsers)
	xload.DecodeJsonnet("onboardings", &gOnboardings)
}

type goldenUsers struct {
	Consumers *variadicUsers `json:"consumers,omitempty"`
}

type variadicUsers struct {
	Foo *User `json:"foo,omitempty"`
	Bar *User `json:"bar,omitempty"`
}

type goldenOnboardings struct {
	Consumers *variadicOnboardings `json:"consumers,omitempty"`
}

type variadicOnboardings struct {
	Foo *Onboarding `json:"foo,omitempty"`
	Bar *Onboarding `json:"bar,omitempty"`
}
