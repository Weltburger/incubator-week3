package models

type Country struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
}
