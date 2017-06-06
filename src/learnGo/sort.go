package main

import (
	"fmt"
)

// 商品参与排序的属性有四个 权重、仓库该商品结算价格、该仓库保存的商品保质期、仓库库存
// 规则；1、若设定了权重，则直接按权重排序，越大越靠前
//       2、未设定权重，按照 结算价、保质期、库存 的优先级进行排序，其中结算价越低越靠前，其他越高越靠前
// 	  3、若a商品价格比b商品低，但b商品的保质期和库存均比a商品高，那b商品排在a前面
// 	  4、任意参与了排序的元素都直接剔除，不参与其他商品的排序

type Goods struct {
	GoodsId   int64       //goodsId
	Weight    interface{} //权重
	Price     int         //价格
	ShelfTime int         //保质期
	Stock     int         //库存
}

func (goods *Goods) sorter() int {
	score := 0
	if goods.Weight != nil {
		return goods.Weight.(int)
	}
	max := 10000000
	score += max - goods.Price
	//todo ...
	return score

}
func main() {

	goods := &Goods{1234, 23, 12, 34, 56}
	fmt.Println(goods)

	goods.sorter()

}
