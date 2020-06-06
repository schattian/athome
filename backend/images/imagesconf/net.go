package imagesconf

import (
	"github.com/athomecomar/envconf"
)

func GetPORT() (port string) {
	switch envconf.GetENV() {
	case envconf.Development:
		port = ":9903"
	case envconf.Staging, envconf.Production:
		port = ":9903"
	}
	return
}
