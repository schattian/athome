package authconf

import (
	"os"

	"github.com/athomecomar/envconf"
)

const DATABASE_SCHEME = "postgres"

func GetDATABASE_PORT() (port string) {
	return ":6479"
}

func GetDATABASE_NUMBER() (dbnum int) {
	return 0
}

func GetDATABASE_HOST() (host string) {
	return "auth_db"
}

func GetDATABASE_PASSWORD() (pwd string) {
	pwd = os.Getenv("REDIS_PASSWORD")
	if pwd == "" && envconf.NotInDevelopment() {
		panic("nil db pwd given")
	}
	return
}
