package dto

type CreateProductReq struct {
	CategoryID  int64  `json:"category_id,omitempty"`
	Price       int64  `json:"price,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Brand       string `json:"brand,omitempty"`
	Image       Image  `json:"image,omitempty"`
}

type Image struct {
	Name string `json:"name,omitempty"`
	Body []byte `json:"body,omitempty"`
}
