package servicestest

import "github.com/athomecomar/athome/backend/services/ent"

// type variadicTokens struct {
// 	Valid   string `json:"valid,omitempty"`
// 	Expired string `json:"expired,omitempty"`
// }

// type variadicUsers struct {
// 	Foo *ent.User `json:"foo,omitempty"`
// 	Bar *ent.User `json:"bar,omitempty"`
// }
type variadicAvailabilities struct {
	Medic  *variationAvailabilities
	Lawyer *variationAvailabilities
}

type variadicEvents struct {
	Medic  *variationEvents
	Lawyer *variationEvents
}

type variadicCalendars struct {
	Medic  *variationCalendars
	Lawyer *variationCalendars
}

type variationCalendars struct {
	A *ent.Calendar
	B *ent.Calendar
	C *ent.Calendar
}

type variationAvailabilities struct {
	First  *subVariationAvailabilities
	Second *subVariationAvailabilities
	Third  *subVariationAvailabilities
}
type subVariationAvailabilities struct {
	A *ent.Availability
	B *ent.Availability
	C *ent.Availability
}

type variationEvents struct {
	First  *subVariationEvents
	Second *subVariationEvents
	Third  *subVariationEvents
}
type subVariationEvents struct {
	A *ent.Event
	B *ent.Event
	C *ent.Event
}

// type variadicOnboardingIdentifications struct {
// 	Foo *ent.OnboardingIdentification `json:"foo,omitempty"`
// 	Bar *ent.OnboardingIdentification `json:"bar,omitempty"`
// }
