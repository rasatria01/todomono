package models

type Todo struct {
	ID     string `json:"id"`
	Todo   string `json:"todo"`
	Status bool   `json:"status"`
}
