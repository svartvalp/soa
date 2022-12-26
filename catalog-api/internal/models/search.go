package models

type (
	FullProductFilters struct {
		IDs                      []int64  `json:"ids,omitempty"`
		Names                    []string `json:"names,omitempty" json:"names,omitempty"`
		CategoryIDs              []int64  `json:"category_ids,omitempty" json:"categoryIDs,omitempty"`
		ProductCharacteristicIDs []int64  `json:"product_characteristic_ids,omitempty" json:"productCharacteristicIDs,omitempty"`
		CharacteristicIDs        []int64  `json:"characteristic_ids,omitempty" json:"characteristicIDs,omitempty"`
	}

	Filter struct {
		Query     string `json:"query,omitempty"`
		CatID     int64  `json:"cat_id,omitempty"`
		PriceFrom int64  `json:"price_from,omitempty"`
		PriceTo   int64  `json:"price_to,omitempty"`
		Brand     string `json:"brand,omitempty"`
	}

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
