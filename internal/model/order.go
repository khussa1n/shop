package model

type Orders struct {
	ID     int64 `db:"id"`
	Number int64 `db:"number"`
}

type GoodsOrders struct {
	GoodID    int64  `db:"good_id"`
	OrderID   string `db:"order_id"`
	GoodCount int64  `db:"good_count"`
}
