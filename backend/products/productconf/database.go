package productconf

import (
	"fmt"
	"os"

	"github.com/athomecomar/envconf"
)

const DATABASE_SCHEME = "postgres"

func GetDATABASE_PORT() (port string) {
	return ":5432"
}

func GetDATABASE_NAME() (db string) {
	switch envconf.GetENV() {
	case envconf.Development:
		db = "product_dev"
	case envconf.Staging:
		db = "product_stg"
	case envconf.Production:
		db = "product_prd"
	}
	return
}

func GetDATABASE_USER() (product string) {
	product = os.Getenv("POSTGRES_USER")
	if product == "postgres" && envconf.NotInDevelopment() {
		panic("default db product given on non-local env")
	}
	return
}

func GetDATABASE_HOST() (host string) {
	return "products_db"
}

func GetDATABASE_PASSWORD() (pwd string) {
	pwd = os.Getenv("POSTGRES_PASSWORD")
	if pwd == "" && envconf.NotInDevelopment() {
		panic("nil db pwd given")
	}
	return
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
