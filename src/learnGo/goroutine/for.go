package main

import (
	"fmt"
	"time"
)

type Url struct {
	u string
}

func main() {

	urls := []string{"a", "b", "c", "d"}

	for _, u := range urls {

		go print(&u)
	}
	urls2 := []*Url{
		newUrl("A"),
		newUrl("B"),
		newUrl("C"),
		newUrl("D"),
	}

	for index, _ := range urls2 {

		go func(i int) {
			fmt.Println(urls2[i].u)
		}(index)
	}
	time.Sleep(2 * time.Second)

}

func newUrl(u string) *Url {

	url := new(Url)
	url.u = u
	return url
}

func print(url *string) {
	//the same pointer
	fmt.Println(url)
}
