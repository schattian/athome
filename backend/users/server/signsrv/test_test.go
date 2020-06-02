package signsrv

import (
	"github.com/athomecomar/athome/backend/users/internal/usertest"
)

var (
	gUsers                     usertest.GoldenUsers
	gOnboardings               usertest.GoldenOnboardings
	gOnboardingIdentifications usertest.GoldenOnboardingIdentifications
	gTokens                    usertest.GoldenTokens
)

func init() {
	usertest.Init(&gTokens, &gUsers, &gOnboardings, &gOnboardingIdentifications)
}
