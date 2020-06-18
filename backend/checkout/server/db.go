package server

import (
	"github.com/athomecomar/athome/backend/checkout/checkoutconf"
	"github.com/athomecomar/storeql/name"
	"github.com/athomecomar/xerrors"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc/status"
)

func ConnDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", checkoutconf.GetDATABASE_SRC())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "sqlx.Open: %v", err)
	}
	db.MapperFunc(name.ToSnakeCase)
	return db, nil
}
