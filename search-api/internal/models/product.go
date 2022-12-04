package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

const ProductTableName = "product"

type Filter struct {
}

type (
	ProductInfo struct {
		ID          int64  `db:"id" json:"id,omitempty"`
		CategoryID  int64  `db:"category_id" json:"category_id,omitempty"`
		Price       int64  `db:"price" json:"price,omitempty"`
		Name        string `db:"name" json:"name,omitempty"`
		Description string `db:"description" json:"description,omitempty"`
		Brand       string `db:"brand" json:"brand,omitempty"`
		Image       string `db:"image" json:"image,omitempty"`

		Characteristics []ProductCharacteristic `json:"characteristics"`
		Categorys       []Category              `json:"categorys"`
	}

	ProductCharacteristic struct {
		ProductId   int64  `db:"product_id" json:"productId,omitempty"`
		ID          int64  `db:"id" json:"ID,omitempty"`
		Name        string `db:"name" json:"name,omitempty"`
		ChType      string `db:"ch_type" json:"chType,omitempty"`
		Description string `db:"description" json:"description,omitempty"`
	}

	Category struct {
		ID          int64  `db:"id" json:"ID,omitempty"`
		Name        string `db:"name" json:"name,omitempty"`
		Description string `db:"description" json:"description,omitempty"`
		ParentID    int64  `db:"parent_id" json:"parentID,omitempty"`
		Level       int64  `db:"level" json:"level,omitempty"`
	}
)

func (p *ProductCharacteristic) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *ProductCharacteristic) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &p)
}