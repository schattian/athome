package field

import (
	"errors"
	"regexp"
)

type Surname string

func (s Surname) Name() string {
	return "surname"
}

func (s Surname) HumanName() string {
	return "Apellido"
}

func (s Surname) Validate() error {
	if s == "" {
		return errors.New("El apellido es obligatorio")
	}

	if len(s) > 30 {
		return errors.New("El apellido no puede tener más de 30 caracteres")
	}
	re := regexp.MustCompile(`^[\w'\-,.][^0-9_!¡?÷?¿/\\+=@#$%ˆ&*(){}|~<>;:[\]]{2,}$`)
	if !re.Match([]byte(s)) {
		return errors.New("Apellido inválido")
	}
	return nil
}
