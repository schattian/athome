package server

import (
	"github.com/athomecomar/athome/backend/users/userconf"
	"github.com/athomecomar/storeql/name"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func connDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", userconf.GetDATABASE_SRC())
	if err != nil {
		return nil, errors.Wrap(err, "sqlx.Open")
	}
	db.MapperFunc(name.ToSnakeCase)
	return db, nil
}
