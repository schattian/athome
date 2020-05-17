package userconf

import (
	"fmt"

	"github.com/athomecomar/envconf"
)

const DATABASE_SCHEME = "postgres"

func GetDATABASE_PORT() (port string) {
	return ":5432"
}

func GetDATABASE_NAME() (db string) {
	switch envconf.GetENV() {
	case envconf.Development:
		db = "user_dev"
	case envconf.Staging:
		db = "user_stg"
	case envconf.Production:
		db = "user_prd"
	}
	return
}

func GetDATABASE_USER() (user string) {
	switch envconf.GetENV() {
	case envconf.Development:
		user = "postgres"
	case envconf.Staging, envconf.Production:
		user = "postgres"
	}
	return
}

func GetDATABASE_HOST() (host string) {
	return "db"
}

func GetDATABASE_PASSWORD() (pwd string) {
	return "secret"
}

func GetDATABASE_SRC() string {
	return fmt.Sprintf(
		"%s://%s:%s@%s%s/%s?sslmode=%s",
		DATABASE_SCHEME, GetDATABASE_USER(), GetDATABASE_PASSWORD(), GetDATABASE_HOST(), GetDATABASE_PORT(),
		GetDATABASE_NAME(), GetDATABASE_SSLMODE(),
	)
}

func GetDATABASE_SSLMODE() (sslmode string) {
	switch envconf.GetENV() {
	case envconf.Development:
		sslmode = "disable"
	case envconf.Staging, envconf.Production:
		sslmode = "require"
	}
	return
}
