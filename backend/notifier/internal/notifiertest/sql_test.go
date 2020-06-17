package notifiertest

import (
	"testing"

	"github.com/athomecomar/athome/backend/notifier/ent"
	"github.com/athomecomar/storeql/test/sqltest"
)

func TestSQL(t *testing.T) {
	sqltest.SQL(t, &ent.Notification{}, "Notification")
}
