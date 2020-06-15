package notifierconf

import (
	"github.com/athomecomar/envconf"
)

func GetNOTIFICATION_JWT_SECRET() string {
	pwd := envconf.Get("NOTIFICATION_JWT_SECRET", "notification_jwt")
	if isSilly(pwd) && envconf.NotInDevelopment() {
		panic("silly notification_jwt secret given")
	}
	return pwd
}

func isSilly(x string) bool {
	return len(x) < 15
}
