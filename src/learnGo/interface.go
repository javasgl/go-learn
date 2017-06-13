package main

import "fmt"

type Goods struct {
	GoodsId string
	Name    string
	Price   float64
	UserId  string
}

func main() {

	goodsList := make([]interface{}, 0, 5)
	goodsList2 := make([]interface{}, 0, 5)
	for i := 0; i < 5; i++ {
		goodsList = append(goodsList, map[string]interface{}{ //cant`t convert to Goods
			"GoodsId": "123445",
			"Name":    "goodsName",
			"Price":   123.5,
			"UserId":  "123",
		})
		goodsList2 = append(goodsList2, &Goods{ //can convert to Goods
			GoodsId: "123445",
			Name:    "goodsName",
			Price:   123.5,
			UserId:  "789",
		})
	}
	for _, goods := range goodsList {

		fmt.Println(goods)
		if goods, ok := goods.(Goods); ok {
			fmt.Println(goods.GoodsId)
			goods.Price += 200
		}
	}

	for _, goods := range goodsList2 {
		fmt.Println(goods)
		if goods, ok := goods.(Goods); ok {
			fmt.Println("is Goods:", goods)
			goods.Price += 120
		}
		fmt.Println(goods)
	}
	fmt.Println(goodsList2)

}
