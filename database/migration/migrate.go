package migration

import (
	"book/database/connect"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func Init() {
	con := connect.GetInstance()

	driver, err := postgres.WithInstance(con.GetDb(), &postgres.Config{})

	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migration/sql",
		"postgres", driver)

	if err != nil {
		panic(err)
	}
	m.Steps(2)
}
