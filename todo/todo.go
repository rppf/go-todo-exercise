package todo

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rppf/go-todo-exercise/database"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title  string `json:"title"`
	Task   string `json:"task"`
	Author string `json:"author"`
}

func GetTodos(c *fiber.Ctx) error {
	db := database.DB
	var todos []Todo
	db.Find(&todos)
	return c.JSON(todos)
}

func GetTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var todo Todo
	db.Find(&todo, id)
	return c.JSON(todo)
}

func NewTodo(c *fiber.Ctx) error {
	db := database.DB

	newTodo := new(Todo)
	if err := c.BodyParser(newTodo); err != nil {
		c.Status(500).Send([]byte(err.Error()))
	}

	db.Create(&newTodo)
	return c.JSON(newTodo)
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var deleteTodo Todo
	db.First(&deleteTodo, id)
	if deleteTodo.Title == "" {
		return c.Status(500).Send([]byte("No book found"))
	}
	db.Delete(&deleteTodo)
	return c.Send([]byte("Book Deleted"))
}
