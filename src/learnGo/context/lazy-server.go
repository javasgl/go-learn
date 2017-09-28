package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", lazyHandler)
	http.ListenAndServe(":9800", nil)
}

func lazyHandler(w http.ResponseWriter, r *http.Request) {
	randNum := rand.Intn(2)

	if randNum == 0 {
		time.Sleep(6 * time.Second)
		fmt.Fprintf(w, "res:slow response.%d\n", randNum)
		fmt.Printf("log:slow response.%d\n", randNum)
		return
	}
	fmt.Fprintf(w, "res:quick response.%d\n", randNum)
	fmt.Printf("log:quick response.%d\n", randNum)
	return
}
