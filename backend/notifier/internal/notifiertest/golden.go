package notifiertest

import "github.com/athomecomar/xtest/xload"

func Init(gPbUsers *GoldenPbUsers, gNotifications *GoldenNotifications) {
	xload.DecodeJsonnet("users", gPbUsers)
	xload.DecodeJsonnet("notifications", gNotifications)
}

type GoldenPbUsers struct {
	Consumers *variadicPbUsers `json:"consumers,omitempty"`
	Merchants *variadicPbUsers `json:"merchants,omitempty"`
}

type GoldenNotifications struct {
	Consumers        *variadicNotifications
	Merchants        *variadicNotifications
	ServiceProviders struct {
		Medic  *variadicNotifications
		Lawyer *variadicNotifications
	}
}
