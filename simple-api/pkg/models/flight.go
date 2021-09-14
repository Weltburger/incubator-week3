package models

import "time"

type Flight struct {
	ID   int       `json:"id" gorm:"primaryKey"`
	From int       `json:"from"`
	To   int       `json:"to"`
	Date time.Time `json:"date"`
}
