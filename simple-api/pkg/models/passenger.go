package models

type Passenger struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
