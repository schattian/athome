package servicestest

import (
	"github.com/athomecomar/athome/backend/services/ent"
	"github.com/athomecomar/xtest/xload"
)

func Init(gServices *GoldenServices, gCalendars *GoldenCalendars, gAvailabilities *GoldenAvailabilities, gEvents *GoldenEvents) {
	xload.DecodeJsonnet("services", gServices)
	xload.DecodeJsonnet("calendars", gCalendars)
	xload.DecodeJsonnet("availabilities", gAvailabilities)
	xload.DecodeJsonnet("events", gEvents)
}

// type GoldenTokens struct {
// 	Sign   *variadicTokens `json:"sign,omitempty"`
// 	Forgot *variadicTokens `json:"forgot,omitempty"`
// }

type GoldenServices struct {
	Foo *ent.Service
	Bar *ent.Service
}

type GoldenCalendars struct {
	Foo *variadicCalendars
	Bar *variadicCalendars
}

type GoldenAvailabilities struct {
	Foo *variadicAvailabilities
	Bar *variadicAvailabilities
}

type GoldenEvents struct {
	Foo *variadicEvents
	Bar *variadicEvents
}

// type GoldenOnboardings struct {
// 	Consumers        *variadicOnboardings `json:"consumers,omitempty"`
// 	ServiceProviders struct {
// 		Medic  *variadicOnboardings `json:"medic,omitempty"`
// 		Lawyer *variadicOnboardings `json:"lawyer,omitempty"`
// 	} `json:"service_providers,omitempty"`
// }

// type GoldenOnboardingIdentifications struct {
// 	ServiceProviders struct {
// 		Medic  *variadicOnboardingIdentifications `json:"medic,omitempty"`
// 		Lawyer *variadicOnboardingIdentifications `json:"lawyer,omitempty"`
// 	} `json:"service_providers,omitempty"`
// }
