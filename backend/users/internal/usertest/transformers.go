package usertest

import (
	"testing"

	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/athome/backend/users/internal/userjwt"
	"github.com/athomecomar/athome/backend/users/pb/pbuser"
	"github.com/pkg/errors"
)

func OnboardingToSignUpSharedRequest(o *ent.Onboarding) *pbuser.SignUpSharedRequest {
	return &pbuser.SignUpSharedRequest{OnboardingId: o.Id, Email: string(o.Email), Name: string(o.Name), Surname: string(o.Surname)}

}

func OnboardingToSignUpStartRequest(o *ent.Onboarding) *pbuser.SignUpStartRequest {
	return &pbuser.SignUpStartRequest{Role: string(o.Role)}
}

func SetStage(o *ent.Onboarding, s field.Stage) *ent.Onboarding {
	o.Stage = s
	return o
}

func OnboardingToSignUpEndRequest(o *ent.Onboarding, pwd string) *pbuser.SignUpEndRequest {
	return &pbuser.SignUpEndRequest{
		OnboardingId: o.Id,
		Password:     pwd,
	}
}

func UserToSignInUserUnsafe(t *testing.T, user *ent.User) *pbuser.SignInUser {
	t.Helper()
	token, err := userjwt.CreateSignToken(user.Id)
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
