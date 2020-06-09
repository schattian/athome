package addresstest

import (
	"testing"

	"github.com/athomecomar/athome/backend/address/ent"
	"github.com/athomecomar/storeql/test/sqltest"
)

func TestSQL(t *testing.T) {
	sqltest.SQL(t, &ent.Address{}, "Address")
}
