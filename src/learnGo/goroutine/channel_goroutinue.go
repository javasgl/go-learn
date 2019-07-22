package main

import (
	"errors"
	"fmt"

	"golang.org/x/sync/errgroup"
)

func main() {

	var eg errgroup.Group

	numChan := make(chan int, 10)

	go func() {
		for n := range numChan {
			fmt.Println("chan----", n)
		}
	}()

	for i := 0; i < 10; i++ {

		func(i int) {

			eg.Go(func() error {
				if i%2 == 0 {

					numChan <- i
				} else {
					fmt.Println("eg---", i)
					return errors.New("xxx")
				}
				return nil
			})
		}(i)
	}
	eg.Wait()
	close(numChan)

	return
}
