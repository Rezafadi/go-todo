package reqres

import "gorm.io/gorm"

type SubTodoRequest struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Status      int      `json:"status"`
	Image       []string `json:"image"`
	TodoID      int      `json:"todo_id"`
}

type SubTodoResponse struct {
	gorm.Model
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Status      int          `json:"status"`
	Image       []string     `json:"image"`
	Todo        GlobalIDName `json:"todo_id"`
}

type GlobalIDName struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
