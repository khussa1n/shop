package dto

import "github.com/khussa1n/shop/internal/model"

type GoodWithOrders struct {
	Good              model.Goods
	OrderNumber       int64
	GoodsCount        int64
	AdditionalShelves []string
}

type AllGoodsByOrders map[string][]GoodWithOrders
