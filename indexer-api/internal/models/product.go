package models

type (
	ProductInfo struct {
		ID          int64  `json:"id,omitempty"`
		CategoryID  int64  `json:"category_id,omitempty"`
		Price       int64  `json:"price,omitempty"`
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
		Brand       string `json:"brand,omitempty"`
		Image       string `json:"image,omitempty"`

		Characteristics []Characteristic `json:"characteristics"`
		Categorys       []Category       `json:"categorys"`
	}

	Category struct {
		ID          int64  `json:"ID,omitempty"`
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
		ParentID    int64  `json:"parentID,omitempty"`
		Level       int64  `json:"level,omitempty"`
	}

	Characteristic struct {
		ID          int64  `json:"ID,omitempty"`
		Name        string `json:"name,omitempty"`
		ChType      string `json:"chType,omitempty"`
		Description string `json:"description,omitempty"`
	}
)
