package field

import "errors"

type Password string

func (p Password) Name() string {
	return "password"
}

func (p Password) HumanName() string {
	return "Contraseña"
}

func (p Password) Validate() error {
	if len(p) < 6 {
		return errors.New("La contraseña debe ser de 6 caracteres como mínimo")
	}
	if len(p) > 25 {
		return errors.New("La contraseña debe ser de 25 caracteres como máximo")
	}
	return nil
}
