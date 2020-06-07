package productconf

import "github.com/athomecomar/envconf"

func GetPORT() (port string) {
	switch envconf.GetENV() {
	case envconf.Development:
		port = ":9992"
	case envconf.Staging, envconf.Production:
		port = ":9992"
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

func GetSEMANTIC_ADDR() (addr string) {
	switch envconf.GetENV() {
	case envconf.Development:
		addr = "semantic_svc:9991"
	}
	return
}

func GetIMAGES_ADDR() (addr string) {
	switch envconf.GetENV() {
	case envconf.Development:
		addr = "images_svc:9993"
	}
	return
}
