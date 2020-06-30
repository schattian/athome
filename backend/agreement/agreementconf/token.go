package agreementconf

import (
	"time"
)

func GetAGREEMENT_TOKEN_EXP() time.Duration {
	return 1 * time.Hour
}
