package models

const ProductTableName = "product"

type Product struct {
	ID          int64  `db:"id"`
	CategoryID  int64  `db:"category_id"`
	Price       int64  `db:"price"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Brand       string `db:"brand"`
	Image       string `db:"image"`
}
