package imageconf

import (
	"github.com/athomecomar/envconf"
)

func GetPORT() (port string) {
	switch envconf.GetENV() {
	case envconf.Development:
		port = ":9993"
	case envconf.Staging, envconf.Production:
		port = ":9993"
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
