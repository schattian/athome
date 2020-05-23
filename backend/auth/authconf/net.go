package authconf

import (
	"time"

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

func GetAUTHENTICATE_JWT_EXP() time.Duration {
	return 5 * time.Hour
}
