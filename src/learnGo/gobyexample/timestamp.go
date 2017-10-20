package main

import (
	"fmt"
	"time"
)

func main() {

	p := fmt.Println
	t := time.Now()

	p(t)
	sec := t.Unix()
	p("秒", t.Unix())
	p("毫秒", t.Unix()*1000)
	p("微秒", t.UnixNano()/1000000)
	p("纳秒", t.UnixNano())

	p(time.Unix(sec, 0))

	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
