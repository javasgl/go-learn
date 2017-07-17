package main

import (
	"io"
	"learnGo/goroutine/pool"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGoroutinues = 25
	polledResource = 2
)

type dbConn struct {
	Id int32
}

func (conn *dbConn) Close() error {
	log.Println("Close dbConn:", conn.Id)
	return nil
}

var idCounter int32

func createConn() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection ", id)
	return &dbConn{id}, nil
}

func performQueries(query int, p *pool.Pool) {
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}
	defer p.Release(conn)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("Query: QueryId[%d],connId[%d]", query, conn.(*dbConn).Id)
}
func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutinues)
	p, err := pool.New(createConn, polledResource)
	if err != nil {
		log.Println(err)
	}
	for query := 0; query <= maxGoroutinues; query++ {
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}
	wg.Wait()
	log.Println("Shutdown....")
	p.Close()
}
