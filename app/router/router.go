package router

import (
	"todo-list/app/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "todo-list/docs"
)

// @title Swagger Todo List
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v1
func Init(app *echo.Echo) {
	app.Use(middleware.CORS())

	app.GET("/swagger/*", echoSwagger.WrapHandler)

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
