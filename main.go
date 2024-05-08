package main

import (
	"log"
	"todo-list/app/router"
	"todo-list/config"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	errVal := godotenv.Load()
	if errVal != nil {
		log.Fatal("Error loading .env file")
	}

	app := echo.New()

	app.Use(middleware.CORS())

	config.Database()

	router.Init(app)

	app.Start(":8080")

	// tambahkan swagger
}
