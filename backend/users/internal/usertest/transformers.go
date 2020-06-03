package usertest

import (
	"testing"

	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/athome/backend/users/internal/userjwt"
	"github.com/athomecomar/athome/backend/users/pb/pbusers"
	"github.com/pkg/errors"
)

func OnboardingToSignUpSharedRequest(o *ent.Onboarding) *pbusers.SignUpSharedRequest {
	return &pbusers.SignUpSharedRequest{OnboardingId: o.Id, Email: string(o.Email), Name: string(o.Name), Surname: string(o.Surname)}

}

func OnboardingToSignUpStartRequest(o *ent.Onboarding) *pbusers.SignUpStartRequest {
	return &pbusers.SignUpStartRequest{Role: string(o.Role)}
}

func SetStage(o *ent.Onboarding, s field.Stage) *ent.Onboarding {
	o.Stage = s
	return o
}

func OnboardingToSignUpEndRequest(o *ent.Onboarding, pwd string) *pbusers.SignUpEndRequest {
	return &pbusers.SignUpEndRequest{
		OnboardingId: o.Id,
		Password:     pwd,
	}
}

func UserToSignInUserUnsafe(t *testing.T, user *ent.User) *pbusers.SignInUser {
	t.Helper()
	token, err := userjwt.CreateSignToken(user.Id)
	if err != nil {
		t.Fatalf(errors.Wrap(err, "CreateSignToken").Error())
	}
	return &pbusers.SignInUser{
		Id:        user.Id,
		SignToken: token,
		Email:     string(user.Email),
		Role:      string(user.Role),
		Name:      string(user.Name),
		Surname:   string(user.Surname),
	}
}
