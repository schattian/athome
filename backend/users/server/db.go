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
	// wd, err := os.Getwd()
	// if err != nil {
	// return nil, errors.Wrap(err, "os.Getwd")
	// }

	// m, err := migrate.New(
	// 	"file://"+wd+"/migrations",
	// 	userconf.GetDATABASE_SRC(),
	// )
	// if err != nil {
	// 	log.Fatal(errors.Wrap(err, "migrate.New"))
	// }
	// if err := m.Up(); err != nil {
	// log.Fatal(errors.Wrap(err, "migrate.Up"))
	// }

	db, err := sqlx.Open("postgres", userconf.GetDATABASE_SRC())

	if err != nil {
		return nil, errors.Wrap(err, "sqlx.Open")
	}
	// driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	// if err != nil {
	// 	return nil, errors.Wrap(err, "postgres.WithInstance")
	// }
	// migrations, err := migrate.NewWithDatabaseInstance(
	// 	"file://"+wd+"/migrations",
	// 	userconf.DATABASE_SCHEME, driver,
	// )
	// if err != nil {
	// 	return nil, errors.Wrap(err, "migrate.NewWithDatabaseInstance")
	// }
	// err = migrations.Steps(1)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "migrations.Steps")
	// }

	db.MapperFunc(name.ToSnakeCase)
	return db, nil
}
