package dto

type CreateCategoryReq struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ParentID    int64  `json:"parent_id,omitempty"`
	Level       int64  `json:"level,omitempty"`
}
