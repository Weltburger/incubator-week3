package storage

import (
	"database/sql"
	_ "github.com/mailru/go-clickhouse"
	"log"
)

type Database struct {
	DB            *sql.DB
	Tx            *sql.Tx
	Stmt          *sql.Stmt
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
	db, err := sql.Open("clickhouse", "http://localhost:8123/trade?&debug=true")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	database.DB = db
}

func (database *Database) CloseDB()  {
	database.DB.Close()
}
