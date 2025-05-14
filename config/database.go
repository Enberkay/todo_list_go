package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"go_todo/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:1111@tcp(127.0.0.1:3306)/go_todo_db?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = database.AutoMigrate(&models.User{}, &models.Todo{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	DB = database
	fmt.Println("Database connected successfully!")
}
