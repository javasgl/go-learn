package main

import (
	"fmt"

	"golang.org/x/sync/errgroup"
)

func main() {
	var (
		eg   = &errgroup.Group{}
		info = &struct {
			a bool
			b bool
			c bool
			d bool
		}{}
	)

	eg.Go(func() error {
		info.a = true
		return nil
	})
	eg.Go(func() error {
		info.b = true
		return nil
	})
	eg.Go(func() error {
		info.c = true
		return nil
	})
	eg.Go(func() error {
		info.d = true
		return nil
	})

	err := eg.Wait()
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Println(info)
}
