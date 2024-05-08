package config

import (
	"fmt"
	"os"
	"strconv"
	"todo-list/app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Database() *gorm.DB {
	connStr := os.Getenv("DB_URL")
	// fmt.Println("DB_URL: ", connStr)
	DB, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	enableDatabaseAutomigration, _ := strconv.ParseBool(os.Getenv("ENABLE_DATABASE_AUTOMIGRATION"))
	if enableDatabaseAutomigration {
		err = DB.AutoMigrate(
			&models.Todo{},
			&models.SubTodo{},
		)

		if err != nil {
			fmt.Println(err)
			panic("Migration Failed")
		}
	}

	fmt.Println("Connected to Database:")

	return DB
}
