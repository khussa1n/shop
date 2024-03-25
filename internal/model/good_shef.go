package model

type GoodsShelves struct {
	GoodID           int64  `db:"good_id"`
	ShelfID          int64  `db:"shelf_id"`
	MainOrAdditional string `db:"main_or_additional"`
}
