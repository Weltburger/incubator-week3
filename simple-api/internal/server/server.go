package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"simple-api/internal/controller"
)

type Server struct {
	Router *echo.Echo
	Controller *controller.Controller
}

func NewServer() *Server {
	server := &Server{
		Router: echo.New(),
		Controller: controller.New(),
	}

	server.Router.Use(middleware.Logger())
	server.Router.Use(middleware.Recover())

	server.Router.GET("/", handleHome)
	apiGroup := server.Router.Group("/api")

	apiGroup.GET("/query1", server.Controller.TicketController().Query1)
	apiGroup.GET("/query2", server.Controller.PassengerController().Query2)
	apiGroup.GET("/query3", server.Controller.PassengerController().Query3)
	apiGroup.GET("/query4", server.Controller.TicketController().Query4)
	apiGroup.GET("/query5", server.Controller.TicketController().Query5)

	return server
}

func handleHome(c echo.Context) error {
	return c.HTML(http.StatusOK, `<h1>The Server is running</h1>`)
}


