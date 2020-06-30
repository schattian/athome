package authconf

import (
	"os"

	"github.com/athomecomar/envconf"
)

func GetDATABASE_PORT() (port string) {
	return ":6379"
}

func GetDATABASE_NUMBER() (dbnum int) {
	return 0
}

func GetDATABASE_HOST() (host string) {
	return "auth_redis"
}

func GetDATABASE_PASSWORD() (pwd string) {
	pwd = os.Getenv("REDIS_PASSWORD")
	if pwd == "" && envconf.NotInDevelopment() {
		panic("nil db pwd given")
	}
	return
}
