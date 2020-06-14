package semantictest

// func Init(gTokens *GoldenTokens, gUsers *GoldenUsers, gOnboardings *GoldenOnboardings, gOnboardingIdentifications *GoldenOnboardingIdentifications) {
// 	xload.DecodeJsonnet("tokens", gTokens)
// 	xload.DecodeJsonnet("users", gUsers)
// 	xload.DecodeJsonnet("onboardings", gOnboardings)
// 	xload.DecodeJsonnet("onboarding_identifications", gOnboardingIdentifications)
// }

type GoldenServiceProviderCategories struct {
	Foo *variadicServiceProviderCategories
}
type GoldenMerchantCategories struct {
	Foo *variadicMerchantCategories
}

type GoldenProductCategories struct {
	Foo *variadicProductCategories
}

// type GoldenUsers struct {
// 	Consumers        *variadicUsers `json:"consumers,omitempty"`
// 	ServiceProviders struct {
// 		Medic  *variadicUsers `json:"medic,omitempty"`
// 		Lawyer *variadicUsers `json:"lawyer,omitempty"`
// 	} `json:"service_providers,omitempty"`
// }

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
