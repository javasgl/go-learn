package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}
func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1

}

func main() {
	mux := sync.Mutex{}
	ua := UserAges{map[string]int{}, mux}
	ua.Add("songgl", 27)
	fmt.Println(ua.Get("songgl"))
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		ua.Add("songgl", rand.Intn(100))
		wg.Done()
	}
	for i := 0; i < 10; i++ {
		fmt.Println(ua.Get("songgl"))
		wg.Done()
	}
	wg.Wait()

}
