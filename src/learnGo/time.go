package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(time.Date(2018, 10, 1, 0, 0, 0, 0, TZ_CST))

	t := time.Now()

	t1, _ := time.Parse("2006.01.02 15:04:05", "0001-01-01 00:00:00")
	fmt.Println(t1.IsZero())
	fmt.Println(t1.IsZero())

	fmt.Println(t.Format("2006.01.02 15:04:05"))

	fmt.Println(NewDay(t))
	effectDate := Day("20180701")
	deadDate := Day("20180901")

	effectTime, _ := effectDate.Time()
	deadTime, _ := deadDate.Time()

	h := deadTime.Sub(effectTime).Minutes()
	fmt.Println(h)
	fmt.Println(effectTime.Location())

}

var (
	TZ_CST *time.Location = time.FixedZone("CST", 8*3600)
)

type Day string // like 20060102

func NewDay(t time.Time) Day {
	return Day(t.In(TZ_CST).Format("20060102"))
}

func (r Day) Time() (time.Time, error) {
	t, err := time.Parse("20060102 -0700", string(r)+" +0800")
	if err != nil {
		return t, err
	}
	t = t.In(TZ_CST)
	return t, err
}
