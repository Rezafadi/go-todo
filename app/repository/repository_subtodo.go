package repository

import (
	"encoding/json"
	"todo-list/app/models"
	"todo-list/app/reqres"
	"todo-list/config"
)

func CreateSubTodo(data reqres.SubTodoRequest) (response models.SubTodo, err error) {
	image, _ := json.Marshal(data.Image)

	response = models.SubTodo{
		Title:       data.Title,
		Description: data.Description,
		Status:      data.Status,
		Image:       string(image),
		TodoID:      data.TodoID,
	}

	err = config.DB.Create(&response).Error
	if err != nil {
		return
	}

	return
}

func BuildSubTodoResponse(data models.SubTodo) (response reqres.SubTodoResponse) {
	response = reqres.SubTodoResponse{
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

	todo, _ := GetTodoByIDPlain(data.TodoID)
	response.Todo = reqres.GlobalIDName{
		ID:   int(todo.ID),
		Name: todo.Title,
	}

	return
}

func GetSubTodos() (responses []reqres.SubTodoResponse, err error) {
	var data []models.SubTodo
	err = config.DB.Find(&data).Error
	if err != nil {
		return
	}

	for _, item := range data {
		responses = append(responses, BuildSubTodoResponse(item))
	}

	return responses, err
}

func GetSubTodoByID(id int) (response reqres.SubTodoResponse, err error) {
	var data models.SubTodo
	err = config.DB.Where("id = ?", id).First(&data).Error
	if err != nil {
		return
	}

	response = BuildSubTodoResponse(data)

	return
}

func GetSubTodoByIDPlain(id int) (response models.SubTodo, err error) {
	err = config.DB.Where("id = ?", id).First(&response).Error
	return
}

func UpdateSubTodoByID(data models.SubTodo) (response models.SubTodo, err error) {
	err = config.DB.Save(&data).Error
	return
}

func DeleteSubTodoByID(data models.SubTodo) (response models.SubTodo, err error) {
	err = config.DB.Delete(&data).Error
	return
}
