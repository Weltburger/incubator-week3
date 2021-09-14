package models

type City struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	CountryID int    `json:"country_id"`
	Name      string `json:"name"`
}
