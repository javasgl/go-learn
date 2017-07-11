package main

import (
	"learnGo/goroutine/runner"
	"log"
	"os"
	"time"
)

const timeout = 3 * time.Second

func main() {
	log.Println("start to working...")
	r := runner.New(timeout)
	r.Add(createTask(), createTask(), createTask())
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Termination due to timeout")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Termination due to interrupt")
			os.Exit(2)
		}
	}
	log.Println("Process end!")

}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processer - Taks #%d", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
