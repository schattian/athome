package servicestest

import (
	"testing"

	"github.com/athomecomar/athome/backend/services/ent"
	"github.com/athomecomar/storeql/test/sqltest"
)

func TestSQL(t *testing.T) {
	sqltest.SQL(t, &ent.Calendar{}, "Calendar")
}
