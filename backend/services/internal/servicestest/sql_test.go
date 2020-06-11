package servicestest

import (
	"testing"

	"github.com/athomecomar/athome/backend/services/ent"
	"github.com/athomecomar/storeql/test/sqltest"
)

func TestSQL(t *testing.T) {
	sqltest.SQL(t, &ent.Calendar{}, "Calendar")
	sqltest.SQL(t, &ent.Event{}, "Event")
	sqltest.SQL(t, &ent.Registry{}, "Registry")
	sqltest.SQL(t, &ent.Service{}, "Service")
	sqltest.SQL(t, &ent.Availability{}, "Availability")
}
