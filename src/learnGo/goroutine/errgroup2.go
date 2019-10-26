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
		err error
		eg  errgroup.Group
	)

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
	for i := range pos {
		index := i
		eg.Go(func() error {
			fmt.Println("deal.....", index, pos[index].Tr, at)
			temp := pos[index]
			temp.Tr = fmt.Sprintf("done---%d", index)
			pos[index] = temp
			return nil
		})
	}
	err = eg.Wait()
	if err != nil {
		fmt.Println(err)
		return errors.New("test")
	}
	for poId := range pos {

		fmt.Println("get tr from ", poId, (pos[poId]).Tr)
	}

	return nil
}
