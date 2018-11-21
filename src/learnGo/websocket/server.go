/**
 * @author songgl
 * @since 2018-11-06 16:22
 */
package main

import (
	"github.com/gorilla/websocket"
	"learnGo/websocket/utils"
	"log"
	"net/http"
	"time"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		wsConn *websocket.Conn
		conn   *utils.Connection
		err    error
		data   []byte
	)
	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}
	if conn, err = utils.InitConnection(wsConn); err != nil {
		goto ERR
	}
	go func() {
		var (
			err error
		)
		for {
			if err = conn.SendMessage([]byte("heartbeat")); err != nil {
				log.Println("send err:", err)
				return
			}
			time.Sleep(time.Second)
		}
	}()
	for {
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		if err = conn.SendMessage(data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()

}
func main() {
	http.HandleFunc("/ws", wsHandler)
	log.Fatal(http.ListenAndServe(":1234", nil))
}
