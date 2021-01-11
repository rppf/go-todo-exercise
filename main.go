package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rppf/go-todo-exercise/database"
	"github.com/rppf/go-todo-exercise/todo"
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/todo", todo.GetTodos)
	app.Get("/api/v1/todo/:id", todo.GetTodo)
	app.Post("/api/v1/todo", todo.NewTodo)
	app.Delete("/api/v1/todo/:id", todo.DeleteTodo)
	app.Put("/api/v1/todo/:id", todo.UpdateTodo)
}

func initDatabase() {
	var err error
	database.DB, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connection to the database.")
	}
	fmt.Println("Database connected")

	database.DB.AutoMigrate(&todo.Todo{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()

	initDatabase()

	setupRoutes(app)

	app.Listen(":3000")
}
