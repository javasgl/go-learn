package main

import (
	"context"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {

	go http.ListenAndServe(":8989", nil)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()

	log.Println(A(ctx))

	select {}

}

func A(ctx context.Context) string {

	go log.Println(B(ctx))
	select {
	case <-ctx.Done():
		return "A done"
	}
	return ""
}

func B(ctx context.Context) string {

	ctx, _ = context.WithCancel(ctx)
	go log.Println(C(ctx))
	select {
	case <-ctx.Done():
		return "B done"
	}
	return ""
}

func C(ctx context.Context) string {
	select {
	case <-ctx.Done():
		return "C done"
	}
	return ""
}
