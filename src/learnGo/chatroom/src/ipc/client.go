package ipc

import (
	"encoding/json"
)

type IpcClient struct {
	conn chan string
}

func NewIpcClient(server *IpcServer) *IpcClient {
	conn := server.Connect()
	return &IpcClient{conn}
}

func (client *IpcClient) Call(method, params string) (response *Response, err error) {

	req := &Request{method, params}

	var request []byte
	request, err = json.Marshal(req)

	if err != nil {
		return
	}
	client.conn <- string(request)
	str := <-client.conn
	var resp Response
	err = json.Unmarshal([]byte(str), &resp)
	response = &resp
	return
}
func (client *IpcClient) Close() {
	client.conn <- "CLOSE"
}
