package main

import (
	"database/sql"
	"flag"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/jinzhu/gorm"
)

var (
	driver string
)

func main() {
	flag.StringVar(&driver, "driver", "", "")
	flag.Parse()

	fmt.Println(driver)

	switch driver {
	case "gorm":
		gormDriver()
	case "xorm":
		xormDriver()
	case "go":
		goDriver()
	}
}

var dsn = "testproxy:testproxy@tcp(10.20.62.36:3333)/information_schema?charset=utf8"

func gormDriver() {
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Exec("select count(1),sleep(30) from PROCESSLIST").Error
	if err != nil {
		fmt.Println("gorm exec failed.", err)
		return
	}
}

func xormDriver() {
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		fmt.Println("xorm: new engine err:", err)
		return
	}
	res, err := engine.NewSession().Exec("select count(2) ,sleep(30) from processlist")
	if err != nil {
		fmt.Println("xorm: exec err:", err)
		return
	}
	fmt.Println("xorm exec success,", fmt.Sprintf("+v", res))
}

func goDriver() {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("go: open err:", err)
		return
	}
	db.SetConnMaxLifetime(10 * time.Minute)

	if MysqlDbErr := db.Ping(); nil != MysqlDbErr {
		panic("go connect failed: " + MysqlDbErr.Error())
	}
	res, err := db.Exec("select count(3) ,sleep(30) from processlist")
	if err != nil {
		fmt.Println("go: exec err:", err)
		return
	}
	fmt.Println("go exec success: ", fmt.Sprintf("%+v", res))
}
