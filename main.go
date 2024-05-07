package main

import (
	"todo-list/app/router"
	"todo-list/config"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	config.Database()

	router.Init(app)

	app.Start(":8080")
}
