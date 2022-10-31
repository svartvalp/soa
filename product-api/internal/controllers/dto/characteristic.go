package dto

type CreateCharacteristicReq struct {
	Name        string `json:"name,omitempty"`
	ChType      string `json:"parent_id,omitempty"`
	Description string `json:"description,omitempty"`
}
