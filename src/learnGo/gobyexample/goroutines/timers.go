package main

import (
	"fmt"
	"time"
)

func main() {

	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C
	fmt.Println("timer1 executed")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer2 executed")
	}()
	stop := timer2.Stop()
	if stop {
		fmt.Println("timer2 stopped")
	}

}
