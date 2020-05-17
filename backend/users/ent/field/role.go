package field

import "errors"

type Role string

func (r Role) Name() string {
	return "role"
}

func (r Role) HumanName() string {
	return "Rol"
}

const (
	ServiceProvider Role = "service-provider"
	Merchant        Role = "merchant"
	Consumer        Role = "consumer"
)

var Roles = []Role{ServiceProvider, Merchant, Consumer}

func (r Role) Validate() error {
	for _, validRole := range Roles {
		if validRole == r {
			return nil
		}
	}
	return errors.New("El rol no es válido")
}
