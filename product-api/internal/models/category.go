package models

const CategoryTableName = "category"

type Category struct {
	ID          int64  `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	ParentID    int64  `db:"parent_id"`
	Level       int64  `db:"level"`
}
