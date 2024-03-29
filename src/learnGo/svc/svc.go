/**
 * @author songgl
 * @since 2018-06-26 10:05
 */
package main

import (
	"fmt"
	"github.com/judwhite/go-svc/svc"
	"log"
	"syscall"
	"time"
)

type program struct {
}

func main() {
	prg := &program{}
	if err := svc.Run(prg, syscall.SIGUSR1, syscall.SIGINT, syscall.SIGTERM); err != nil {
		log.Fatal(err)
	}
}

func (program) Init(env svc.Environment) error {

	fmt.Println("init.....")
	return nil
}

func (program) Start() error {

	fmt.Println("start.....")
	fmt.Println(syscall.Getpid())
	go func() {
		ticker := time.NewTicker(2 * time.Second)
		for t := range ticker.C {
			fmt.Println("tick at", t)
		}
	}()
	return nil
}
func (program) Stop() error {

	fmt.Println("stop.....")
	return nil
}
