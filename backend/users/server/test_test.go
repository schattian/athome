package server

import (
	"testing"

	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/athome/backend/users/pbuser"
	"github.com/athomecomar/xtest/xload"
	"github.com/pkg/errors"
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

func onboardingToSignUpSharedRequest(o *ent.Onboarding) *pbuser.SignUpSharedRequest {
	return &pbuser.SignUpSharedRequest{OnboardingId: o.Id, Email: string(o.Email), Name: string(o.Name), Surname: string(o.Surname)}

}

func onboardingToSignUpStartRequest(o *ent.Onboarding) *pbuser.SignUpStartRequest {
	return &pbuser.SignUpStartRequest{Role: string(o.Role)}
}

func setStage(o *ent.Onboarding, s field.Stage) *ent.Onboarding {
	o.Stage = s
	return o
}
func onboardingToSignUpEndRequest(o *ent.Onboarding, pwd string) *pbuser.SignUpEndRequest {
	return &pbuser.SignUpEndRequest{
		OnboardingId: o.Id,
		Password:     pwd,
	}
}

func userToSignInUserUnsafe(t *testing.T, user *ent.User) *pbuser.SignInUser {
	t.Helper()
	token, err := createSignToken(user.Id)
	if err != nil {
		t.Fatalf(errors.Wrap(err, "CreateSignToken").Error())
	}
	return &pbuser.SignInUser{
		Id:        user.Id,
		SignToken: token,
		Email:     string(user.Email),
		Role:      string(user.Role),
		Name:      string(user.Name),
		Surname:   string(user.Surname),
	}
}
