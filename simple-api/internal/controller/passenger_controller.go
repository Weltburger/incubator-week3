package controller

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"simple-api/pkg/models"
	"strconv"
)

type PassengerController struct {
	controller *Controller
}

func (passengerController *PassengerController) Query2(c echo.Context) error {
	passengers := make([]models.Passenger, 0)
	ctx := context.Background()
	passengers = passengerController.controller.DB.PassengerRepository().Query2(ctx)

	return c.JSON(http.StatusOK, passengers)
}

func (passengerController *PassengerController) Query3(c echo.Context) error {
	var count int
	ctx := context.Background()
	count = passengerController.controller.DB.PassengerRepository().Query3(ctx)

	return c.String(http.StatusOK, "Количество пассажиров, которые заказывали больше одного билета: " + strconv.Itoa(count))
}
