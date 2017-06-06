package calc

import (
	"fmt"
	"net/http"
	"time"

	"github.com/hprose/hprose-golang/rpc"
)

func hello(name string) string {
	return "Hello " + name + "!"
}
func talk2(msg string) string {
	return fmt.Sprintf("%s(%d)", msg, time.Now().Unix())
}

func main() {
	service := rpc.NewHTTPService()
	service.AddFunction("hello", hello, rpc.Options{})
	service.AddFunction("talk", talk2, rpc.Options{})
	http.ListenAndServe(":8111", service)
}
