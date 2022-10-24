package models

const CharacteristicTableName = "characteristic"

type Characteristic struct {
	ID           int64  `db:"id"`
	Name         string `db:"name"`
	ChType       string `db:"ch_type"`
	Descriptions string `db:"descriptions"`
}
