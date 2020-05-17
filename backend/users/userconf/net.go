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
