package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string

type User struct {
	Name String
	Age  int
	Sex  bool
}

func (str String) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, str)
}
func (user User) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, user.Name, user.Age, user.Sex)
}
func main() {
	http.Client
	http.Transport
	http.Handle("/string", String("I am a string!"))
	http.Handle("/user", User{"songgl", 20, false})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
