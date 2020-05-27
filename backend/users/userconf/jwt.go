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
	pwd := "sign_jwt"
	if env := os.Getenv("SIGN_JWT_SECRET"); env != "" {
		pwd = env
	}
	if isSilly(pwd) && envconf.NotInDevelopment() {
		panic("silly sign_jwt secret given")
	}
	return pwd
}

func GetFORGOT_JWT_EXP() time.Duration {
	return 10 * time.Hour
}

func GetFORGOT_JWT_SECRET() string {
	pwd := "forgot_jwt"
	if env := os.Getenv("FORGOT_JWT_SECRET"); env != "" {
		pwd = env
	}
	if isSilly(pwd) && envconf.NotInDevelopment() {
		panic("silly forgot_jwt secret given")
	}
	return pwd
}

func isSilly(x string) bool {
	return len(x) < 10
}
