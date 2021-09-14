package models

type Ticket struct {
	ID          int     `json:"id" gorm:"primaryKey"`
	PassengerID int     `json:"name"`
	FlightID    int     `json:"flight_id"`
	ClassID     int     `json:"class_id"`
	LunchID     int     `json:"lunch_id"`
	Price       float32 `json:"price"`
	Weight      int     `json:"weight"`
}
