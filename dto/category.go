package dto

type Category struct {
	Id    int64  `db:"category_id" json:"category_id"`
	Title string `db:"category_title" json:"category_title"`
}
