package authconf

import (
	"os"

	"github.com/athomecomar/envconf"
)

func GetAUTHENTICATE_JWT_SECRET() string {
	pwd := os.Getenv("AUTHENTICATE_JWT_SECRET")
	if isSilly(pwd) && envconf.NotInDevelopment() {
		panic("silly jwt given")
	}
	return pwd
}

func isSilly(x string) bool {
	return len(x) < 10
}
