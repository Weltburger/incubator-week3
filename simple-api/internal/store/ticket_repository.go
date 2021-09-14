package store

import (
	"context"
	"simple-api/pkg/models"
)

type TicketRepository struct {
	database *Database
}

func (ticketRepository *TicketRepository) Query1(ctx context.Context) []models.Ticket {
	var tickets []models.Ticket
	ticketRepository.database.GormDB.WithContext(ctx).
		Raw(`SELECT tt."id", tt.passenger_id, tt.flight_id, tt.class_id, tt.lunch_id, tt.price, tt.weight
					FROM ticket tt
					JOIN (SELECT ft."id" 
						  FROM flight ft
						  JOIN (SELECT cy."id" 
								FROM city cy
								WHERE cy."name"='Berlin') cy ON cy."id" = ft."to"
						 WHERE ft."date" BETWEEN '2021-08-12' AND '2021-09-12') ft ON ft.id = tt.flight_id`).
		Scan(&tickets)

	return tickets
}

func (ticketRepository *TicketRepository) Query4(ctx context.Context) int {
	var mass int
	ticketRepository.database.GormDB.WithContext(ctx).
		Raw(`SELECT SUM(tt.weight) 
				 FROM ticket tt
				 JOIN (SELECT ft."id" 
					   FROM flight ft
					   WHERE ft."date" BETWEEN '2021-08-12' AND '2021-09-12') ft ON ft."id" = tt.flight_id`).
		Scan(&mass)

	return mass
}

func (ticketRepository *TicketRepository) Query5(ctx context.Context) float32 {
	var avg float32
	ticketRepository.database.GormDB.WithContext(ctx).
		Raw(`SELECT AVG(weight) AS average_weight FROM ticket;`).Scan(&avg)

	return avg
}
