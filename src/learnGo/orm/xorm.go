package main

import (
	"fmt"
	"time"

	"github.com/fatih/structs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	engine *xorm.Engine
)

type TestOrder struct {
	Id         int64
	Fee        float64    `xorm:"'fee'"`
	AFee       float64    `xorm:"actually_fee"`
	SubOrders  []SubOrder `xorm:"-"`
	Version    int        `xorm:"version 'version'"`
	CreateTime int        `xorm:"created 'create_time'"`
	UpdateTime int        `xorm:"updated 'update_time'"`
	PayTime    time.Time  `xorm:"NOT NULL 'pay_time'"`
	// PayTime        time.Time `xorm:"'pay_time'"`
	SettlementMode int `xorm:"'settlement_mode' not null default 0"`
}

type SubOrder struct {
	Id int64
}

func main() {
	fmt.Println(new(time.Time))
	fmt.Println(new(time.Time).Format(time.RFC3339))
	fmt.Println(new(time.Time).Format("0001-01-01 00:00:00 +0000 UTC"))

	engine, err := xorm.NewEngine("mysql", "root:@/test?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	engine.ShowSQL(true)

	// engine.DropTables(&TestOrder{})
	// engine.CreateTables(&TestOrder{})
	_, err = engine.Insert(&TestOrder{Fee: 0.666, AFee: 0.8888})
	fmt.Println("INSERT:", err)

	//AllCols
	//MustCols
	//Cols
	//坑
	var to TestOrder
	engine.Id(1).Get(&to)
	fmt.Println(fmt.Sprintf("%#v", to))
	fmt.Println(to.PayTime.Format("2018-09-12 14:34:23"))
	fmt.Println(to.PayTime.Format("0001-01-01 00:00:00 +0000 UTC"))
	fmt.Println(to.PayTime.Format(time.RFC3339))
	// engine.Id(1).AllCols().Update(&TestOrder{Id: 1, AFee: 0.12})
	to.AFee += 0.11112
	to.Fee = 0.0
	_, err = engine.Id(1).Update(&to)
	fmt.Println("UPDATE:", err)

	m := structs.Map(to)
	m["fee"] = 0.1
	fmt.Println(m)
	//engine.Id(1).Update(m)
}
