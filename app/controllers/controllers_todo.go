package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-list/app/repository"
	"todo-list/app/reqres"

	"github.com/labstack/echo/v4"
)

// CreateTodo godoc
// @Summary Create Todo
// @Description Create Todo
// @Tags Todo
// @Accept json
// @Produce json
// @Param Todo body reqres.TodoRequest true "Create Todo"
// @Success 200
// @Router /v1/todo [post]
func CreateTodo(c echo.Context) error {
	var input reqres.TodoRequest
	if err := c.Bind(&input); err != nil {
		return err
	}

	data, err := repository.CreateTodo(input)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": 200,
		"data":    data,
		"message": "success create todo",
	})
}

// GetTodos godoc
// @Summary Get Todos
// @Description Get Todos
// @Tags Todo
// @Accept json
// @Produce json
// @Success 200
// @Router /v1/todo [get]
func GetTodos(c echo.Context) error {
	data, err := repository.GetTodos()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": 200,
		"data":    data,
		"message": "success get todos",
	})
}

// GetTodoByID godoc
// @Summary Get Todo By ID
// @Description Get Todo By ID
// @Tags Todo
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200
// @Router /v1/todo/{id} [get]
func GetTodoByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := repository.GetTodoByID(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": 200,
		"data":    data,
		"message": "success get todo by id",
	})
}

// UpdateTodoByID godoc
// @Summary Update Todo By ID
// @Description Update Todo By ID
// @Tags Todo
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Param Todo body reqres.TodoRequest true "Update Todo"
// @Success 200
// @Router /v1/todo/{id} [put]
func UpdateTodoByID(c echo.Context) error {
	var input reqres.TodoRequest
	if err := c.Bind(&input); err != nil {
		return err
	}

	image, _ := json.Marshal(input.Image)

	id, _ := strconv.Atoi(c.Param("id"))
	data, err := repository.GetTodoByIDPlain(id)
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

	update, err := repository.UpdateTodoByID(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Error update data")
	}

	dataUpdate, _ := repository.GetTodoByID(int(update.ID))

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": 200,
		"data":    dataUpdate,
		"message": "success update todo by id",
	})
}

// DeleteTodoByID godoc
// @Summary Delete Todo By ID
// @Description Delete Todo By ID
// @Tags Todo
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200
// @Router /v1/todo/{id} [delete]
func DeleteTodoByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := repository.GetTodoByIDPlain(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "data not found")
	}

	_, err = repository.DeleteTodoByID(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Error delete data")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": 200,
		"data":    data,
		"message": "success delete todo by id",
	})
}
