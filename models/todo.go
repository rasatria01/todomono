package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Todo   string `json:"todo"`
	Status string `json:"status"`
}

func (todo *Todo) TableName() string {
	return "t_todo"
}
