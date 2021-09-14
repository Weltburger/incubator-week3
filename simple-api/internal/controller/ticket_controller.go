package controller

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"simple-api/pkg/models"
	"strconv"
)

type TicketController struct {
	controller *Controller
}

func (ticketController *TicketController) Query1(c echo.Context) error {
	tickets := make([]models.Ticket, 0)
	ctx := context.Background()
	tickets = ticketController.controller.DB.TicketRepository().Query1(ctx)

	return c.JSON(http.StatusOK, tickets)
}

func (ticketController *TicketController) Query4(c echo.Context) error {
	var mass int
	ctx := context.Background()
	mass = ticketController.controller.DB.TicketRepository().Query4(ctx)

	return c.String(http.StatusOK, "Общая масса перевезенного багажа за месяц: " + strconv.Itoa(mass))
}

func (ticketController *TicketController) Query5(c echo.Context) error {
	var avg float32
	ctx := context.Background()
	avg = ticketController.controller.DB.TicketRepository().Query5(ctx)

	return c.String(http.StatusOK, "Средний багаж пассажиров: " + fmt.Sprintf("%f", avg))
}
