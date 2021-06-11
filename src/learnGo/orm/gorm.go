package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type TestOrderV2 struct {
	gorm.Model
	Fee        float64
	AFee       float64
	Version    int
	CreateTime int64
	UpdateTime int64
}

func main() {

	db, err := gorm.Open("mysql", "root@/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&TestOrderV2{})
	// db.Create(&TestOrderV2{Fee: 12.34, AFee: 23.12, Version: 1, CreateTime: time.Now().Unix()})
	var order TestOrderV2
	db.Debug().Unscoped().First(&order, 1)
	db.Debug().Unscoped().Model(&order).Update("Fee", 2.84)
	db.Debug().Unscoped().Model(&order).Update("AFee", 0)
	order.Version = 1
	db.Debug().Unscoped().Model(&order).Update()
}
