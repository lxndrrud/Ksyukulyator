package dto

import (
	"database/sql"
	"strconv"
)

type ProductInterface interface {
}

type Product struct {
	Id         int64         `db:"product_id" json:"product_id"`
	Title      string        `db:"product_title" json:"product_title"`
	Cost       float32       `db:"product_cost" json:"product_cost"`
	IdCategory sql.NullInt64 `db:"product_id_category"`
}

type ProductFull struct {
	Product
	Category
}

func (p Product) GetString() (string, string) {
	return p.Title, strconv.FormatFloat(float64(p.Cost), 'f', 2, 32)
}
