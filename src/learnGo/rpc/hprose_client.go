package calc

import (
	"fmt"

	"github.com/hprose/hprose-golang/rpc"
)

type Stub struct {
	Hello      func(string) (string, error)
	AsyncHello func(func(string, error), string) `name:"hello"`
	Talk       func(string) (string, error)
	AsyncTalk  func(func(string, error), string) `name:"talk"`
}

func talk() {
	client := rpc.NewClient("http://127.0.0.1:8111/")
	var stub *Stub

	client.UseService(&stub)
	stub.AsyncTalk(func(res string, err error) {
		fmt.Println(res, err)
	}, "this is async msg from hprose client")

	fmt.Println(stub.Talk("this is msg  from hprose client"))
}

// func main() {
// 	client := rpc.NewClient("http://127.0.0.1:8111/")
// 	var stub *Stub
// 	client.UseService(&stub)
// 	stub.AsyncHello(func(result string, err error) {
// 		fmt.Println(result, err)
// 	}, "async world")
// 	fmt.Println(stub.Hello("world"))

// 	stub.AsyncTalk(func(res string, err error) {
// 		fmt.Println(res, err)
// 	}, "this is async msg from hprose client")

// 	fmt.Println(stub.Talk("this is msg  from hprose client"))
// }
