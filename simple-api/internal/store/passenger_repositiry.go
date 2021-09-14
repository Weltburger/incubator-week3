package store

import (
	"context"
	"simple-api/pkg/models"
)

type PassengerRepository struct {
	database *Database
}

func (passengerRepository *PassengerRepository) Query2(ctx context.Context) []models.Passenger {
	var passengers []models.Passenger
	passengerRepository.database.GormDB.WithContext(ctx).
		Raw(`SELECT DISTINCT pr."id", pr."name", pr.age, pr.email, pr.phone_number 
				 FROM passenger as pr
				 JOIN ticket as tt ON tt.passenger_id = pr."id"
				 WHERE tt.weight > 10`).
		Scan(&passengers)

	return passengers
}

func (passengerRepository *PassengerRepository) Query3(ctx context.Context) int {
	var count int
	passengerRepository.database.GormDB.WithContext(ctx).
		Raw(`SELECT COUNT(*) FROM (SELECT COUNT(tt.passenger_id)
			 	 FROM ticket as tt
				 GROUP BY tt.passenger_id
				 HAVING COUNT(tt.passenger_id) > 1) as res`).
		Scan(&count)

	return count
}
