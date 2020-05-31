package email

import (
	"net/mail"

	"github.com/matcornic/hermes/v2"
)

func ForgotPasswordHeaders(from, to mail.Address) map[string][]string {
	return map[string][]string{
		"From":    {from.String()},
		"To":      {to.String()},
		"Subject": {"Recuperación de contraseña"},
	}
}

func ForgotPassword(name string, tokenByRole map[string]string) hermes.Email {
	base := hermes.Email{
		Body: hermes.Body{
			Name:     name,
			Greeting: "Hola",
			Intros: []string{
				"Has recibido este mensaje porque se ha solicitado recuperar la contraseña.",
			},
			Outros: []string{
				"Si no has solicitado el cambio de contraseña, ignora este mensaje.",
			},
			Signature: "Muchas gracias",
		},
	}
	var actions []hermes.Action

	for role, token := range tokenByRole {
		action := hermes.Action{
			// Instructions: fmt.Sprintf("Haz click en el siguiente botón para restablecer tu contraseña como %v:", RoleTranslations[role]),
			Button: hermes.Button{
				Color: ColorByRole[role],
				Text:  "Cambiar contraseña de " + RoleTranslations[role],
				Link:  "https://athome.com.ar/reset-password?token=" + token,
			},
		}
		actions = append(actions, action)
	}
	if len(actions) == 1 {
		actions[0].Button.Text = "Cambiar contraseña"
	}
	base.Body.Actions = actions
	return base
}
