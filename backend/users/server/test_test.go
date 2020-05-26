package server

import (
	"testing"

	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/athome/backend/users/pbuser"
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
	Foo *ent.User `json:"foo,omitempty"`
	Bar *ent.User `json:"bar,omitempty"`
}

type goldenOnboardings struct {
	Consumers *variadicOnboardings `json:"consumers,omitempty"`
}

type variadicOnboardings struct {
	Foo *ent.Onboarding `json:"foo,omitempty"`
	Bar *ent.Onboarding `json:"bar,omitempty"`
}

func onboardingToSignUpSharedRequest(t *testing.T, o *ent.Onboarding) *pbuser.SignUpSharedRequest {
	t.Helper()
	return &pbuser.SignUpSharedRequest{OnboardingId: o.Id, Email: string(o.Email), Name: string(o.Name), Surname: string(o.Surname)}

}

func onboardingToSignUpStartRequest(t *testing.T, o *ent.Onboarding) *pbuser.SignUpStartRequest {
	t.Helper()
	return &pbuser.SignUpStartRequest{Role: string(o.Role)}
}

func setStage(o *ent.Onboarding, s field.Stage) *ent.Onboarding {
	o.Stage = s
	return o
}
