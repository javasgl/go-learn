package main

import (
	"learnGo/goroutine/work"
	"log"
	"runtime"
	"sync"
	"time"
)

var names = []string{
	"steve", "bob", "mary", "jason", "therese",
}

type namePrinter struct {
	name string
}

func (m *namePrinter) Task() {
	log.Println(m.name, runtime.NumGoroutine())
	time.Sleep(1 * time.Millisecond)
}

func main() {
	p := work.New(2)
	var wg sync.WaitGroup
	wg.Add(100 * len(names))
	for i := 0; i < 100; i++ {
		for _, name := range names {
			np := namePrinter{
				name: name,
			}
			go func() {
				p.Run(&np)
				wg.Done()
			}()
		}
	}
	wg.Wait()
	log.Println(runtime.NumGoroutine())
	p.Shutdown()
}
