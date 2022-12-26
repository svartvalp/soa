package models

const ProductTableName = "product"

type Product struct {
	ID          int64  `db:"id" json:"id,omitempty"`
	CategoryID  int64  `db:"category_id" json:"category_id,omitempty"`
	Price       int64  `db:"price" json:"price,omitempty"`
	Name        string `db:"name" json:"name,omitempty"`
	Description string `db:"description" json:"description,omitempty"`
	Brand       string `db:"brand" json:"brand,omitempty"`
	Image       string `db:"image" json:"image,omitempty"`
}

type ProductFilters struct {
	IDs                      []int64  `json:"ids,omitempty"`
	Names                    []string `json:"names,omitempty" json:"names,omitempty"`
	CategoryIDs              []int64  `json:"category_ids,omitempty" json:"categoryIDs,omitempty"`
	ProductCharacteristicIDs []int64  `json:"product_characteristic_ids,omitempty" json:"productCharacteristicIDs,omitempty"`
	CharacteristicIDs        []int64  `json:"characteristic_ids,omitempty" json:"characteristicIDs,omitempty"`
}

type FullProductInfo struct {
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
