package userconf

import "github.com/athomecomar/envconf"

func GetPORT() (port string) {
	switch envconf.GetENV() {
	case envconf.Development:
		port = ":9990"
	case envconf.Staging, envconf.Production:
		port = ":9990"
	}
	return
}

func GetAUTH_ADDR() (addr string) {
	switch envconf.GetENV() {
	case envconf.Development:
		addr = "auth_svc:9900"
	}
	return
}

func GetMAILER_ADDR() (addr string) {
	switch envconf.GetENV() {
	case envconf.Development:
		addr = "mailer_svc:9901"
	}
	return
}
