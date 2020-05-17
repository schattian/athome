package field

import (
	"errors"
	"regexp"
)

type Name string

func (n Name) Name() string {
	return "name"
}

func (n Name) HumanName() string {
	return "Nombre"
}

func (n Name) Validate() error {
	if n == "" {
		return errors.New("El nombre es obligatorio")
	}
	if len(n) > 30 {
		return errors.New("El nombre no puede tener más de 30 caracteres")
	}
	re := regexp.MustCompile(`^[\w'\-,.][^0-9_!¡?÷?¿/\\+=@#$%ˆ&*(){}|~<>;:[\]]{2,}$`)
	if !re.Match([]byte(n)) {
		return errors.New("Nombre inválido")
	}
	return nil
}
