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
	app.Delete("/api/v1/book/:id", todo.DeleteTodo)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connection to the database.")
	}
	fmt.Println("Database connected")
}

func main() {
	app := fiber.New()

	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(":3000")
}
