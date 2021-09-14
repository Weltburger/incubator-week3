package models

type Class struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
}
