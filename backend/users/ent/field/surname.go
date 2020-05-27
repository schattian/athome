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
		return errors.New("surname must exist")
	}

	if len(s) > 30 {
		return errors.New("surname len cant be > 30")
	}
	re := regexp.MustCompile(`^[\w'\-,.][^0-9_!¡?÷?¿/\\+=@#$%ˆ&*(){}|~<>;:[\]]{2,}$`)
	if !re.Match([]byte(s)) {
		return errors.New("invalid surname")
	}
	return nil
}
