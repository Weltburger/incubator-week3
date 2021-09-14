package controller

import (
	"simple-api/internal/store"
)

type Controller struct {
	DB *store.Database
	ticketController *TicketController
	passengerController *PassengerController
}

func (controller *Controller) TicketController() *TicketController {
	if controller.ticketController != nil {
		return controller.ticketController
	}

	controller.ticketController = &TicketController{controller: controller}

	return controller.ticketController
}

func (controller *Controller) PassengerController() *PassengerController {
	if controller.passengerController != nil {
		return controller.passengerController
	}

	controller.passengerController = &PassengerController{controller: controller}

	return controller.passengerController
}

func New() *Controller {
	db := store.GetDB()
	db.Migrate()
	return &Controller{DB: db}
}
