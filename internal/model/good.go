package model

type Goods struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}
