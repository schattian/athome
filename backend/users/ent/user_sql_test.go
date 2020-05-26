package ent

import (
	"testing"

	"github.com/athomecomar/storeql/test/sqltest"
)

func TestUserSQL(t *testing.T) {
	sqltest.SQL(t, gUsers.Consumers.Foo, "User")
}
