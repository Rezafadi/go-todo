package reqres

import "gorm.io/gorm"

type TodoRequest struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Status      int      `json:"status"`
	Image       []string `json:"image"`
}

type TodoResponse struct {
	gorm.Model
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Status      int      `json:"status"`
	Image       []string `json:"image"`
}
