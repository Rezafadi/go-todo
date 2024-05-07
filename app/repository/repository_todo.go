package repository

import (
	"encoding/json"
	"todo-list/app/models"
	"todo-list/app/reqres"
	"todo-list/config"
)

func CreateTodo(data reqres.TodoRequest) (response models.Todo, err error) {
	image, _ := json.Marshal(data.Image)

	response = models.Todo{
		Title:       data.Title,
		Description: data.Description,
		Status:      data.Status,
		Image:       string(image),
	}

	err = config.DB.Create(&response).Error
	if err != nil {
		return
	}

	return
}

func BuildTodoResponse(data models.Todo) (response reqres.TodoResponse) {
	response = reqres.TodoResponse{
		Model:       data.Model,
		Title:       data.Title,
		Description: data.Description,
		Status:      data.Status,
	}

	var image []string
	err := json.Unmarshal([]byte(data.Image), &image)
	if err != nil {
		return
	}
	response.Image = image

	return
}

func GetTodos() (responses []reqres.TodoResponse, err error) {
	var data []models.Todo
	err = config.DB.Find(&data).Error
	if err != nil {
		return
	}

	for _, item := range data {
		responses = append(responses, BuildTodoResponse(item))
	}

	return responses, err
}

func GetTodoByID(id int) (response reqres.TodoResponse, err error) {
	var data models.Todo
	err = config.DB.Where("id = ?", id).First(&data).Error
	if err != nil {
		return
	}

	response = BuildTodoResponse(data)

	return
}

func GetTodoByIDPlain(id int) (response models.Todo, err error) {
	err = config.DB.Where("id = ?", id).First(&response).Error
	return
}

func UpdateTodoByID(data models.Todo) (response models.Todo, err error) {
	err = config.DB.Save(&data).Error
	return
}

func DeleteTodoByID(data models.Todo) (response models.Todo, err error) {
	err = config.DB.Delete(&data).Error
	return
}
