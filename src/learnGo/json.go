package main

import (
	"encoding/json"
	"fmt"
)

type Msg struct {
	code int
	data string
	time int64
}

func main() {
	var m Msg
	m.code = 1
	m.data = "json data"
	m.time = 12919212212
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println(b)
	} else {
		fmt.Println(err)
	}

	data := make([]map[string]string, 2)

	data[0] = map[string]string{

		"title":   "this is title",
		"content": "this is content",
	}

	b, err = json.Marshal(data)
	fmt.Println(string(b))

}
