package userconf

import (
	"os"
	"time"

	"github.com/athomecomar/envconf"
)

func GetSIGN_JWT_EXP() time.Duration {
	return 2 * time.Minute
}

func GetSIGN_JWT_SECRET() string {
	pwd := os.Getenv("SIGN_JWT_SECRET")
	if isSilly(pwd) && envconf.NotInDevelopment() {
		panic("silly sign_jwt secret given")
	}
	return pwd
}

func isSilly(x string) bool {
	return len(x) < 10
}
