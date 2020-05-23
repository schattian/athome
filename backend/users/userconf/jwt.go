package userconf

import (
	"os"
	"time"

	"github.com/athomecomar/envconf"
)

func GetSIGN_JWT_EXP() time.Duration {
	return 5 * time.Second
}

func GetSIGN_JWT_SECRET() string {
	pwd := os.Getenv("SIGN_JWT_SECRET")
	if isSilly(pwd) && envconf.NotInDevelopment() {
		panic("nil db pwd given")
	}
	return pwd
}

func isSilly(x string) bool {
	return len(x) < 10
}
