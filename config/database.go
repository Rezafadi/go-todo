package config

import (
	"fmt"
	"todo-list/app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Database() *gorm.DB {
	connStr := "postgresql://todo-app_owner:PCh3e4OFWnmU@ep-blue-frog-a1gn1oeo.ap-southeast-1.aws.neon.tech/todo-app?sslmode=require"
	DB, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	enableDatabaseAutomigration := false
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
