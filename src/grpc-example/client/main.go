/**
 * @author songgl
 * @since 2018-11-23 16:40
 */
package main

import (
	"log"
	"os"

	"grpc-example/hello"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"qiniu.io/order-rule/engine/sdk"
)

const (
	address     = "localhost:50001"
	defaultName = "songgl"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {

		log.Fatalf("did not connect :%v", err)
	}
	defer conn.Close()

	ruleClient := engine.NewRuleEngineClient(conn)
	ruleClient.ExcuteRule(context.Background(), &engine.Rule{})

	c := hello.NewGreeterClient(conn)
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &hello.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
