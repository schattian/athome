package authconf

import (
	"os"
	"time"

	"github.com/athomecomar/envconf"
)

func GetAUTH_JWT_SECRET() string {
	pwd := "auth_jwt"
	if env := os.Getenv("AUTH_JWT_SECRET"); env != "" {
		pwd = env
	}
	if isSilly(pwd) && envconf.NotInDevelopment() {
		panic("silly auth_jwt secret given")
	}
	return pwd
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

func isSilly(x string) bool {
	return len(x) < 10
}

func GetAUTH_JWT_EXP() time.Duration {
	return 5 * time.Hour
}

func GetAUTH_JWT_REFRESH_EXP() time.Duration {
	return 24 * 10 * time.Hour
}
