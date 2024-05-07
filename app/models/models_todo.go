package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title       string `json:"title" gorm:"varchar(255)"`
	Description string `json:"description" gorm:"varchar(255)"`
	Status      int    `json:"status" gorm:"int8"`
	Image       string `json:"image" gorm:"text"`
}

type SubTodo struct {
	gorm.Model
	Title       string `json:"title" gorm:"varchar(255)"`
	Description string `json:"description" gorm:"varchar(255)"`
	Status      int    `json:"status" gorm:"int8"`
	Image       string `json:"image" gorm:"text"`
	TodoID      int    `json:"todo_id" gorm:"int8"`
}
