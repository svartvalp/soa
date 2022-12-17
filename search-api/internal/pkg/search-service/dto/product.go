package dto

type (
	ProductInfo struct {
		ID          string    `db:"id" json:"id,omitempty"`
		CategoryID  float64   `db:"category_id" json:"category_id,omitempty"`
		Price       float64   `db:"price" json:"price,omitempty"`
		Name        string    `db:"name" json:"name,omitempty"`
		Description string    `db:"description" json:"description,omitempty"`
		Brand       string    `db:"brand" json:"brand,omitempty"`
		Image       string    `db:"image" json:"image,omitempty"`
		CategoryIDs []float64 `json:"category_ids"`

		Characteristics []ProductCharacteristic `json:"characteristics"`
		Categorys       []Category              `json:"categorys"`
	}

	ProductCharacteristic struct {
		ProductId   float64 `db:"product_id" json:"productId,omitempty"`
		ID          float64 `db:"id" json:"ID,omitempty"`
		Name        string  `db:"name" json:"name,omitempty"`
		ChType      string  `db:"ch_type" json:"chType,omitempty"`
		Description string  `db:"description" json:"description,omitempty"`
	}

	Category struct {
		ID          float64 `db:"id" json:"ID,omitempty"`
		Name        string  `db:"name" json:"name,omitempty"`
		Description string  `db:"description" json:"description,omitempty"`
		ParentID    float64 `db:"parent_id" json:"parentID,omitempty"`
		Level       float64 `db:"level" json:"level,omitempty"`
	}
)
