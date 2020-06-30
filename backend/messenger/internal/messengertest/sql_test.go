package messengertest

import (
	"testing"

	"github.com/athomecomar/storeql/test/sqltest"
)

func TestSQL(t *testing.T) {
	sqltest.SQL(t, gMessages.Foo, "Messages")
}
