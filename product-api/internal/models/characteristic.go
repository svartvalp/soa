package models

const (
	ProductCharacteristicTableName = "product_characteristic"
	CharacteristicTableName        = "characteristic"
)

type (
	Characteristic struct {
		ID          int64  `db:"id" json:"ID,omitempty"`
		Name        string `db:"name" json:"name,omitempty"`
		ChType      string `db:"ch_type" json:"chType,omitempty"`
		Description string `db:"description" json:"description,omitempty"`
	}

	ProductCharacteristic struct {
		ProductId   int64  `db:"product_id" json:"productId,omitempty"`
		ID          int64  `db:"id" json:"ID,omitempty"`
		Name        string `db:"name" json:"name,omitempty"`
		ChType      string `db:"ch_type" json:"chType,omitempty"`
		Description string `db:"description" json:"description,omitempty"`
	}
)
