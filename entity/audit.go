package entity

type (
	Audit struct {
		ID           string `json:"_key"`
		CreatedAt    string `json:"createdAt"`
		UpdatedAt    string `json:"updatedAt"`
		DeletedAt    string `json:"deletedAt"`
		CurrNo       int    `json:"currNo"`
		InputterName string `json:"inputterName"`
	}
)
