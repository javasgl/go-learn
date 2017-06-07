package ipc

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Method string "method"
	Params string "params"
}

type Response struct {
	Code string "code"
	Body string "body"
}

type Server interface {
	Name() string
	Handle(method, params string) *Response
}

type IpcServer struct {
	Server
}

func NewIpcServer(server Server) *IpcServer {
	return &IpcServer{server}
}

func (ipcServer *IpcServer) Connect() chan string {

	session := make(chan string, 0)
	go func(ch chan string) {
		for {
			request := <-ch
			if request == "CLOSE" {
				break
			}
			var req Request
			err := json.Unmarshal([]byte(request), &req)
			if err != nil {
				fmt.Println("Invalid request format:", request)
			}
			resp := ipcServer.Handle(req.Method, req.Params)
			response, err := json.Marshal(resp)
			ch <- string(response)
		}
		fmt.Println("Session Closed.")
	}(session)

	fmt.Println("A New Session has been created successfully.")
	return session
}
