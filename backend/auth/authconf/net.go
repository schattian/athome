package authconf

import (
	"github.com/athomecomar/envconf"
)

func GetPORT() (port string) {
	switch envconf.GetENV() {
	case envconf.Development:
		port = ":9990"
	case envconf.Staging, envconf.Production:
		port = ":30050"
	}
	return
}
