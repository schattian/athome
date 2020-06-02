package field

import "errors"

type DNI uint64

func (d DNI) Name() string {
	return "dni"
}

func (d DNI) HumanName() string {
	return "DNI"
}

func (d DNI) Validate() error {
	if d < 1e6 || d > 1e8 {
		return errors.New("DNI inválido")
	}
	return nil
}
