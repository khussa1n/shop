package model

type GoodsOrders struct {
	OrderID   int64 `db:"order_id"`
	GoodID    int64 `db:"good_id"`
	GoodCount int64 `db:"good_count"`
}
