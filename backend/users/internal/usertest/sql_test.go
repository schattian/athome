package usertest

import (
	"testing"

	"github.com/athomecomar/storeql/test/sqltest"
)

func TestOnboardingIdentificationsSQL(t *testing.T) {
	sqltest.SQL(t, gOnboardings.Consumers.Foo, "Onboarding")
	sqltest.SQL(t, gOnboardingIdentifications.ServiceProviders.Medic.Foo, "Onboarding")
	sqltest.SQL(t, gOnboardings.ServiceProviders.Medic.Foo, "Onboarding")
}
