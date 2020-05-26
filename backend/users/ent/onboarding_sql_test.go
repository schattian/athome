package ent

import (
	"testing"

	"github.com/athomecomar/storeql/test/sqltest"
)

func TestOnboardingSQL(t *testing.T) {
	sqltest.SQL(t, gOnboardings.Consumers.Foo, "Onboarding")
}
