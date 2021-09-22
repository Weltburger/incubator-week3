package main

import (
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

type Trade struct {
	E  string `json:"e"`
	E1 int    `json:"E"`
	S  string `json:"s"`
	A  int    `json:"a"`
	P  string `json:"p"`
	Q  string `json:"q"`
	F  int    `json:"f"`
	L  int    `json:"l"`
	T  int    `json:"T"`
	M  bool   `json:"m"`
}

type Request struct {
	Method string    `json:"method"`
	Params [1]string `json:"params"`
	ID     int       `json:"id"`
}

var upgrader = websocket.Upgrader{}

func handler(writer http.ResponseWriter, request *http.Request) {
	var conn, err = upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Fatal(err)
	}

	ws, _, err := websocket.DefaultDialer.Dial("wss://stream.binance.com:9443/ws", nil)
	if err != nil {
		log.Fatal(err)
	}

	data := new(Trade)
	go func(ws *websocket.Conn) {
		for {
			_, message, readErr := ws.ReadMessage()
			if readErr != nil {
				fmt.Println(readErr)
				return
			}

			if err := json.Unmarshal(message, &data); err != nil {
				log.Fatal(err)
			}
			if data.E == "" {
				continue
			}
			tm := strconv.Itoa(data.T)
			s, _ := strconv.ParseInt(tm[:10], 10, 64)
			ns, _ := strconv.ParseInt(tm[10:], 10, 64)

			fmt.Println(data, time.Unix(s, ns))
		}
	}(ws)
	
	go func(conn *websocket.Conn, ws *websocket.Conn) {
		for {
			_, mes, err := conn.ReadMessage()
			if err != nil {
				log.Fatal(err)
			}
			if strings.ToLower(string(mes)) == "stop" {
				os.Exit(1)
			}

			params := strings.Split(string(mes), " ")
			if len(params) > 1 {
				if params[0] == "SUBSCRIBE" {
					newRequest := Request{params[0], [1]string{params[1]+"@trade"}, 1 }
					err = ws.WriteJSON(newRequest)
					if err != nil {
						log.Fatal(err)
					}
				} else if params[0] == "UNSUBSCRIBE"{
					newRequest := Request{params[0], [1]string{params[1]+"@trade"}, 312}
					err = ws.WriteJSON(newRequest)
					if err != nil {
						log.Fatal(err)
					}
				}

			}

		}
	}(conn, ws)
}

func stopper(writer http.ResponseWriter, request *http.Request) {
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

func main() {
	http.HandleFunc("/ws", handler)
	http.HandleFunc("/stop", stopper)

	http.ListenAndServe(":3000", nil)
}
