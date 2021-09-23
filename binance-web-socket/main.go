package main

import (
	"binance-web-socket/api"
	"net/http"
)

func main() {
	api.DB.StartDB()
	defer api.DB.CloseDB()

	http.HandleFunc("/ws", api.Handler)
	http.HandleFunc("/stop", api.Stopper)

	http.ListenAndServe(":3000", nil)
}
