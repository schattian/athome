package mailerconf

import (
	"github.com/athomecomar/envconf"
)

func GetPORT() (port string) {
	switch envconf.GetENV() {
	case envconf.Development:
		port = ":9988"
	case envconf.Staging, envconf.Production:
		port = ":9988"
	}
	return
}
