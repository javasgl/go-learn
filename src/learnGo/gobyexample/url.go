package main

import (
	"fmt"
	"net/url"
)

func main() {

	s := "mysql://user:pass@host.com:3306/database?encoding=utf8#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)
	fmt.Println(u.User)

	fmt.Println(u.Hostname())
	fmt.Println(u.Port())

	fmt.Println(u.Path)
	fmt.Println(u.RawQuery)
	fmt.Println(u.Fragment)

	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
}
