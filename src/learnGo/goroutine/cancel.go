package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		ch := make(chan bool)

		go func() {
			time.Sleep(5 * time.Second)
			ch <- true
		}()
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		case <-ch:
			fmt.Println("Done after 10 s")
		}

	}(ctx)
	wg.Wait()

}
