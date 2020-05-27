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
		return errors.New("password len must be >= 6")
	}
	if len(p) > 25 {
		return errors.New("password len must be <= 25")
	}
	return nil
}
