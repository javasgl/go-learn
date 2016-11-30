package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	age  int32 //private param will not include in the json serialize
}
type Gril struct {
	Perion Person
	Sexy   bool
}

func main() {
	defer func() {
		fmt.Println("program end~")
	}()
	person := Person{"songgl", 18}
	gril := Gril{person, true}
	str, err := json.Marshal(gril)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(str))

}




