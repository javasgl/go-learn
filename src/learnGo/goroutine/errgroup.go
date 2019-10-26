package main

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	err := do()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("do other tings")

}

type Refund struct {
	Tr string
}

func do() error {

	var (
		err    error
		eg     errgroup.Group
		trChan chan string
	)

	trChan = make(chan string, 2)
	pos := make(map[int]Refund, 2)

	pos[3] = Refund{
		Tr: "tr_id_1",
	}
	pos[5] = Refund{
		Tr: "tr_id_2",
	}
	for poId := range pos {
		fmt.Println("source:", poId, pos[poId])
	}

	at := time.Now()
	fmt.Println(pos)
	for i := range pos {
		index := i
		eg.Go(func() error {
			fmt.Println("deal.....", index, pos[index].Tr, at)
			if index == 1 {
				trChan <- pos[index].Tr
			} else {
				return errors.New("failed")
			}
			return nil
		})
	}
	err = eg.Wait()
	if err != nil {
		fmt.Println(err)
		return errors.New("test")
	}
	if len(trChan) != 2 {
		return errors.New("recharge failed")
	}
	close(trChan)
	for tr := range trChan {

		fmt.Println("get tr from ", tr)
	}

	return nil
}
