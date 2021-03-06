package api

import (
	"binance-web-socket/models"
	"binance-web-socket/storage"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	UPGRADER    = websocket.Upgrader{}
	DB          = new(storage.Database)
	STOP        = "stop"
	BINANCEWS   = "wss://stream.binance.com:9443/ws"
	SUBSCRIBE   = "SUBSCRIBE"
	UNSUBSCRIBE = "UNSUBSCRIBE"
	TRADE       = "@trade"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	var conn, err = UPGRADER.Upgrade(writer, request, nil)
	if err != nil {
		log.Fatal(err)
	}

	ws, _, err := websocket.DefaultDialer.Dial(BINANCEWS, nil)
	if err != nil {
		log.Fatal(err)
	}

	data := new(models.Trade)
	go func(ws *websocket.Conn) {
		for {
			_, message, readErr := ws.ReadMessage()
			if readErr != nil {
				fmt.Println(readErr)
				return
			}

			if err := json.Unmarshal(message, data); err != nil {
				log.Fatal(err)
			}
			if data.EventType == "" {
				continue
			}
			tm := strconv.Itoa(data.TradeTime)
			s, _ := strconv.ParseInt(tm[:10], 10, 64)
			ns, _ := strconv.ParseInt(tm[10:], 10, 64)

			ctx := context.Background()
			DB.TradesRepository().AddTrade(ctx, data)

			fmt.Println(data, time.Unix(s, ns))
		}
	}(ws)

	go func(conn *websocket.Conn, ws *websocket.Conn) {
		for {
			_, mes, err := conn.ReadMessage()
			if err != nil {
				log.Fatal(err)
			}
			if strings.ToLower(string(mes)) == STOP {
				os.Exit(1)
			}

			params := strings.Split(string(mes), " ")
			if len(params) > 1 {
				if params[0] == SUBSCRIBE {
					newRequest := models.Request{Method: params[0], Params: [1]string{params[1] + TRADE}, ID: 1}
					err = ws.WriteJSON(newRequest)
					if err != nil {
						log.Fatal(err)
					}
				} else if params[0] == UNSUBSCRIBE {
					newRequest := models.Request{Method: params[0], Params: [1]string{params[1] + TRADE}, ID: 312}
					err = ws.WriteJSON(newRequest)
					if err != nil {
						log.Fatal(err)
					}
				}

			}

		}
	}(conn, ws)
}

func Stopper(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Stopping the application", string(body))
	_, err = writer.Write([]byte("Stopping the application"))
	if err != nil {
		return
	}

	os.Exit(1)
}
