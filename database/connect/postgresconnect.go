package connect

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "postgres"
)

var singleInstance *PostgresConnect

type PostgresConnect struct {
	db   *sql.DB
	look *sync.Mutex
}

func (pc PostgresConnect) GetDb() *sql.DB {
	return pc.db
}

func createConnect() *PostgresConnect {
	db := openDatabase()
	pc := &PostgresConnect{
		db:   db,
		look: &sync.Mutex{},
	}

	return pc
}

func openDatabase() *sql.DB {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable search_path=public",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	return db
}

func GetInstance() *PostgresConnect {
	if singleInstance == nil {
		pc := createConnect()
		pc.look.Lock()
		defer pc.look.Unlock()

		if singleInstance == nil {
			singleInstance = pc
			fmt.Println("Creating single instance now.")
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}
