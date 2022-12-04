package models

const CategoryTableName = "category"

type Category struct {
	ID          int64  `db:"id" json:"ID,omitempty"`
	Name        string `db:"name" json:"name,omitempty"`
	Description string `db:"description" json:"description,omitempty"`
	ParentID    int64  `db:"parent_id" json:"parentID,omitempty"`
	Level       int64  `db:"level" json:"level,omitempty"`
}
