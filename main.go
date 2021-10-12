package main

import (
	"book/database/connect"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	pc := connect.InitPostgresConnect()
	driver, err := postgres.WithInstance(pc.Db, &postgres.Config{})
	fmt.Println(err)
	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		"postgres", driver)
	m.Steps(2)
	fmt.Println(err)
}
