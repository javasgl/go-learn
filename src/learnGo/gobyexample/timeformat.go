package main

import (
	"fmt"
	"time"
)

//time.Format(layout string)
//基于模板格式化时间
//time.Format("2012-12-15 14:02:34")
func main() {
	p := fmt.Println

	t := time.Now()
	p(t)
	p(t.Format(time.RFC3339))
	p(t.Unix())
	p(t.UnixNano())

	t1, _ := time.Parse(time.RFC3339, "2017-10-20T09:15:19+08:00")
	p("parse:", t1)
	p(t.Format("3:00PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	p(t.Format("2006-01-02 15:04:05"))
	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM")
	p(t2)

	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d 00:00\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e := time.Parse(ansic, "8:41PM")
	p(e)

}
