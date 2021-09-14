package models

type Lunch struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
}
