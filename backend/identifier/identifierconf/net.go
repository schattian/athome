package identifierconf

import (
	"github.com/athomecomar/envconf"
)

func GetPORT() (port string) {
	switch envconf.GetENV() {
	case envconf.Development:
		port = ":9989"
	case envconf.Staging, envconf.Production:
		port = ":30056"
	}
	return
}
