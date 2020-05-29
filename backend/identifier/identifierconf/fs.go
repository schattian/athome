package identifierconf

import "github.com/athomecomar/envconf"

func GetDATA_DIR() (dir string) {
	switch envconf.GetENV() {
	case envconf.Production, envconf.Staging:
		dir = "/data"
	default:
		dir = "./data"
	}
	return
}
