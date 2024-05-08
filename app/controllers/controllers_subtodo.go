package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-list/app/repository"
	"todo-list/app/reqres"

	"github.com/labstack/echo/v4"
)

// CreateSubTodo godoc
// @Summary Create SubTodo
// @Description Create SubTodo
// @Tags SubTodo
// @Accept json
// @Produce json
// @Param SubTodo body reqres.SubTodoRequest true "Create SubTodo"
// @Success 200
// @Router /v1/subtodo [post]
func CreateSubTodo(c echo.Context) error {
	var input reqres.SubTodoRequest
	if err := c.Bind(&input); err != nil {
		return err
	}

	data, err := repository.CreateSubTodo(input)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": 200,
		"data":    data,
		"message": "success create todo",
	})
}

// GetSubTodos godoc
// @Summary Get SubTodos
// @Description Get SubTodos
// @Tags SubTodo
// @Accept json
// @Produce json
// @Success 200
// @Router /v1/subtodo [get]
func GetSubTodos(c echo.Context) error {
	data, err := repository.GetSubTodos()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": 200,
		"data":    data,
		"message": "success get todos",
	})
}

// GetSubTodoByID godoc
// @Summary Get SubTodoByID
// @Description Get SubTodoByID
// @Tags SubTodo
// @Accept json
// @Produce json
// @Success 200
// @Router /v1/subtodo/{id} [get]
func GetSubTodoByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := repository.GetSubTodoByID(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": 200,
		"data":    data,
		"message": "success get todo by id",
	})
}

// UpdateSubTodoByID godoc
// @Summary Update SubTodoByID
// @Description Update SubTodoByID
// @Tags SubTodo
// @Accept json
// @Produce json
// @Param id path int true "SubTodo ID"
// @Param SubTodo body reqres.SubTodoRequest true "Update SubTodo"
// @Success 200
// @Router /v1/subtodo/{id} [put]
func UpdateSubTodoByID(c echo.Context) error {
	var input reqres.SubTodoRequest
	if err := c.Bind(&input); err != nil {
		return err
	}

	image, _ := json.Marshal(input.Image)

	id, _ := strconv.Atoi(c.Param("id"))
	data, err := repository.GetSubTodoByIDPlain(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "data not found")
	}

	if input.Title != "" {
		data.Title = input.Title
	}
	if input.Description != "" {
		data.Description = input.Description
	}
	if input.Status != 0 {
		data.Status = input.Status
	}
	if input.Image != nil {
		data.Image = string(image)
	}
	if input.TodoID != 0 {
		data.TodoID = input.TodoID
	}

	update, err := repository.UpdateSubTodoByID(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Error update data")
	}

	dataUpdate, _ := repository.GetSubTodoByID(int(update.ID))

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": 200,
		"data":    dataUpdate,
		"message": "success update todo by id",
	})
}

// DeleteSubTodoByID godoc
// @Summary Delete SubTodoByID
// @Description Delete SubTodoByID
// @Tags SubTodo
// @Accept json
// @Produce json
// @Param id path int true "SubTodo ID"
// @Success 200
// @Router /v1/subtodo/{id} [delete]
func DeleteSubTodoByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := repository.GetSubTodoByIDPlain(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "data not found")
	}

	_, err = repository.DeleteSubTodoByID(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Error delete data")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": 200,
		"data":    data,
		"message": "success delete todo by id",
	})
}
