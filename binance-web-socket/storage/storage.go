package storage

import (
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/mailru/go-clickhouse"
)

type Database struct {
	DB            *redis.Client
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
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	database.DB = client
}

func (database *Database) CloseDB()  {
	database.DB.Close()
}
