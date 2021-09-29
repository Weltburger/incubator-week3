package storage

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type TradesStorage struct {
	database *Database
}

func (tradesStorage *TradesStorage) Set(data []byte) {
	err := tradesStorage.database.DB.Set(uuid.New().String(), data, time.Minute * 2).Err()
	if err != nil {
		fmt.Println(err)
	}
}
