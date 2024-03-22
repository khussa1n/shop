package model

type OrdersByGoods struct {
	OrderNumber int64 `db:"number"`
	GoodID      int64 `db:"good_id"`
	GoodCount   int64 `db:"good_count"`
}

type ShelvesByGoods struct {
	GoodID            int64    `db:"good_id"`
	MainShelf         string   `db:"main_shelf"`
	AdditionalShelves []string `db:"additional_shelves"`
}
