package pbconf

import "github.com/athomecomar/envconf"

type service struct {
	Name string
}

func (s *service) GetAddr() string {
	return s.GetHost() + s.GetPort()
}

func (s *service) GetPort() (port string) {
	switch envconf.GetENV() {
	case envconf.Development:
		port = devPorts[s.Name]
	}
	return
}

func (s *service) GetHost() (h string) {
	switch envconf.GetENV() {
	case envconf.Development:
		h = s.Name + "_svc"
	}
	return
}

var (
	Auth       = &service{Name: "auth"}
	Mailer     = &service{Name: "mailer"}
	Identifier = &service{Name: "identifier"}
	Users      = &service{Name: "users"}
	Notifier   = &service{Name: "notifier"}
	Messager   = &service{Name: "messager"}

	Semantic  = &service{Name: "semantic"}
	Products  = &service{Name: "products"}
	Images    = &service{Name: "images"}
	Services  = &service{Name: "services"}
	Addresses = &service{Name: "address"}
)

var devPorts = map[string]string{
	"auth":       ":9900",
	"mailer":     ":9901",
	"identifier": ":9902",

	"users":    ":9990",
	"semantic": ":9991",
	"products": ":9992",
	"images":   ":9993",
	"services": ":9994",
	"address":  ":9995",
	"notifier": ":9996",
	"messager": ":9997",
}
