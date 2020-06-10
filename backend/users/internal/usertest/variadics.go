package usertest

import "github.com/athomecomar/athome/backend/users/ent"

type variadicTokens struct {
	Valid   string `json:"valid,omitempty"`
	Expired string `json:"expired,omitempty"`
}

type variadicUsers struct {
	Foo *ent.User `json:"foo,omitempty"`
	// Bar *ent.User `json:"bar,omitempty"`
}

type variadicOnboardings struct {
	Foo *ent.Onboarding `json:"foo,omitempty"`
	Bar *ent.Onboarding `json:"bar,omitempty"`
}

type variadicOnboardingIdentifications struct {
	Foo *ent.OnboardingIdentification `json:"foo,omitempty"`
	Bar *ent.OnboardingIdentification `json:"bar,omitempty"`
}
