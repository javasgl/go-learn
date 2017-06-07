package ipc

import "testing"

type EchoServer struct {
}

func (echoServer *EchoServer) Handle(method, params string) *Response {
	resp := &Response{method, params}
	return resp
}

func (echoServer EchoServer) Name() string {
	return "EchoServer"
}
func TestIpc(t *testing.T) {

	server := NewIpcServer(&EchoServer{})

	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	defer client1.Close()
	defer client2.Close()

	resp1, _ := client1.Call("clent1 call", "from client1")
	resp2, _ := client2.Call("clent2 call", "from client2")
	if resp1.Code != "clent1 call" || resp1.Body != "from client1" {
		t.Error("faild")
	}
	if resp2.Code != "clent2 call" || resp2.Body != "from client2" {
		t.Error("faild")
	}
}
