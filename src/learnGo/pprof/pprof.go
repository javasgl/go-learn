package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/test", handler)
	http.ListenAndServe(":9800", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	log.Println(r.Form)

	doSomeThing(10000)

	buff := getSomeBytes()

	b, err := ioutil.ReadAll(buff)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(b)
}

func doSomeThing(times int) {
	for i := 0; i <= times; i++ {
		for j := 0; j <= times; j++ {

		}
	}
}

func getSomeBytes() *bytes.Buffer {
	var buffer bytes.Buffer

	for i := 0; i <= 200000; i++ {
		buffer.Write([]byte{'0' + byte(rand.Intn(10))})
	}

	return &buffer
}
