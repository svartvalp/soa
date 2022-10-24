package models

const ProductCharacteristicTableName = "product_characteristic"

type ProductCharacteristic struct {
	ID        int64  `db:"id"`
	ProductID int64  `db:"product_id"`
	Value     string `db:"value"`
	ChType    string `db:"ch_type"`
}
