package main

import (
	"log"
	"os"
	"strconv"

	"github.com/athomecomar/athome/backend/semantic/semanticconf"
	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"github.com/pkg/errors"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("action not given (up|down|drop|force)")
	}
	action := os.Args[1]

	var err error
	var version int
	if len(os.Args) > 2 {
		version, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("%v version is not a number", os.Args[2])
		}
	}

	m, err := migrate.New(
		"file://migrations",
		semanticconf.GetDATABASE_SRC(),
	)
	if err != nil {
		log.Fatal(errors.Wrap(err, "migrate.New"))
	}
	switch action {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	case "drop":
		err = m.Drop()
	case "force":
		err = m.Force(version)
	}

	if err != nil {
		log.Fatal(errors.Wrap(err, "migrate."+action))
	}
}
