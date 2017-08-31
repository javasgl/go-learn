package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

func main() {
	http.Handle("/", websocket.Handler(Echo))
	if err := http.ListenAndServe(":54321", nil); err != nil {
		log.Fatal("ListenAndServe error:", err)
	}
}

func Echo(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {

			fmt.Println("can`t recevied")
			break
		}

		fmt.Println("recevied from client:", reply)

		msg := "recevied! " + reply + time.Now().String()

		fmt.Println("send to client:", msg)

		if err = websocket.Message.Send(ws, msg); err != nil {

			fmt.Println("can`t send")
			break
		}
		select {

		case <-time.After(time.Second * 5):
			fmt.Println("timer auto send msg to clent")
			websocket.Message.Send(ws, "time.after 5 sec ,send msg!")
		}

		tick := time.NewTicker(time.Second * 3)
		for {

			select {
			case <-tick.C:
				fmt.Println("ticker auto send msg to clent")
				websocket.Message.Send(ws, "time.ticker send msg!")
			}
		}

	}
}
