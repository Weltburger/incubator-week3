package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Database struct {
	db *sql.DB
	tradesStorage *TradesStorage
}

func (database *Database) TradesRepository() *TradesStorage {
	if database.tradesStorage != nil {
		return database.tradesStorage
	}

	database.tradesStorage = &TradesStorage{database: database}

	return database.tradesStorage
}

func (database *Database) StartDB() {
	db, err := sql.Open("postgres", "postgres://postgres:1234@localhost:5432/trades?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	database.db = db
}

func (database *Database) CloseDB()  {
	database.db.Close()
}
