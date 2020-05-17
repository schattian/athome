package field

import (
	"errors"
	"regexp"
)

type Email string

func (e Email) Name() string {
	return "email"
}

func (e Email) HumanName() string {
	return "Correo electrónico"
}

func (e Email) Validate() error {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !re.Match([]byte(e)) {
		return errors.New("Email inválido")
	}
	return nil
}
