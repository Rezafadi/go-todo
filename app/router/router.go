package router

import (
	"todo-list/app/controllers"
	"todo-list/app/middlewares"

	"github.com/labstack/echo/v4"
)

func Init(app *echo.Echo) {
	app.Use(middlewares.Cors())

	v1 := app.Group("/v1")
	{
		todo := v1.Group("/todo")
		{
			todo.GET("", controllers.GetTodos)
			todo.GET("/:id", controllers.GetTodoByID)
			todo.POST("", controllers.CreateTodo)
			todo.PUT("/:id", controllers.UpdateTodoByID)
			todo.DELETE("/:id", controllers.DeleteTodoByID)
		}

		subTodo := v1.Group("/subtodo")
		{
			subTodo.GET("", controllers.GetSubTodos)
			subTodo.GET("/:id", controllers.GetSubTodoByID)
			subTodo.POST("", controllers.CreateSubTodo)
			subTodo.PUT("/:id", controllers.UpdateSubTodoByID)
			subTodo.DELETE("/:id", controllers.DeleteSubTodoByID)
		}
	}
}
