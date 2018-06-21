package main

import (
	"fmt"

	"github.com/fatih/structs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	engine *xorm.Engine
)

type TestOrder struct {
	Id        int64
	Fee       float64    `xorm:"'fee'"`
	AFee      float64    `xorm:"actually_fee"`
	SubOrders []SubOrder `xorm:"-"`
	Version   int        `xorm:"version 'version'"`
}

type SubOrder struct {
	Id int64
}

func main() {
	engine, err := xorm.NewEngine("mysql", "root:@/test?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}

	// engine.DropTables(&TestOrder{})
	engine.CreateTables(&TestOrder{})
	// engine.Insert(&TestOrder{Fee: 0.666, AFee: 0.8888})

	//AllCols
	//MustCols
	//Cols
	//坑
	var to TestOrder
	engine.Id(1).Get(&to)
	fmt.Println(to)
	// engine.Id(1).AllCols().Update(&TestOrder{Id: 1, AFee: 0.12})
	to.AFee += 0.11112
	to.Fee = 0.0
	m := structs.Map(to)
	m["fee"] = 0.1
	fmt.Println(m)

	// engine.Id(1).AllCols().Update(&to)
	engine.Id(1).Update(m)
}
