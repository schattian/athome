package userconf

import "github.com/athomecomar/envconf"

func GetPORT() (port string) {
	switch envconf.GetENV() {
	case envconf.Development:
		port = ":9991"
	case envconf.Staging, envconf.Production:
		port = ":30051"
	}
	return
}

func GetAUTH_ADDR() (addr string) {
	switch envconf.GetENV() {
	case envconf.Development:
		addr = "auth_svc:9990"
	}
	return
}

func GetMAILER_ADDR() (addr string) {
	switch envconf.GetENV() {
	case envconf.Development:
		addr = "auth_svc:9992"
	}
	return
}
