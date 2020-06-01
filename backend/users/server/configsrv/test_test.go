package configsrv

import (
	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/xtest/xload"
)

var (
	gUsers       *goldenUsers
	gOnboardings *goldenOnboardings
	gTokens      *goldenTokens
)

func init() {
	xload.DecodeJsonnet("tokens", &gTokens)
	xload.DecodeJsonnet("users", &gUsers)
	xload.DecodeJsonnet("onboardings", &gOnboardings)
}

type goldenTokens struct {
	Sign   *variadicTokens `json:"sign,omitempty"`
	Forgot *variadicTokens `json:"forgot,omitempty"`
}

type variadicTokens struct {
	Valid   string `json:"valid,omitempty"`
	Expired string `json:"expired,omitempty"`
}

type goldenUsers struct {
	Consumers        *variadicUsers `json:"consumers,omitempty"`
	ServiceProviders *variadicUsers `json:"service_providers,omitempty"`
}

type variadicUsers struct {
	Foo *ent.User `json:"foo,omitempty"`
	Bar *ent.User `json:"bar,omitempty"`
}

type goldenOnboardings struct {
	Consumers        *variadicOnboardings `json:"consumers,omitempty"`
	ServiceProviders *variadicOnboardings `json:"service_providers,omitempty"`
}

type variadicOnboardings struct {
	Foo *ent.Onboarding `json:"foo,omitempty"`
	Bar *ent.Onboarding `json:"bar,omitempty"`
}
