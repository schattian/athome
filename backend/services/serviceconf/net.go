package serviceconf

import "github.com/athomecomar/envconf"

func GetPORT() (port string) {
	switch envconf.GetENV() {
	case envconf.Development:
		port = ":9994"
	case envconf.Staging, envconf.Production:
		port = ":9994"
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

func GetUSERS_ADDR() (addr string) {
	switch envconf.GetENV() {
	case envconf.Development:
		addr = "users_svc:9990"
	}
	return
}

func GetADDRESS_ADDR() (addr string) {
	switch envconf.GetENV() {
	case envconf.Development:
		addr = "address_svc:9995"
	}
	return
}
