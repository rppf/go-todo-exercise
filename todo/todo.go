package todo

import "github.com/gofiber/fiber/v2"

func GetTodos(c *fiber.Ctx) error {
	return c.Send([]byte("All Books"))
}

func GetTodo(c *fiber.Ctx) error {
	return c.Send([]byte("Single Book"))
}

func NewTodo(c *fiber.Ctx) error {
	return c.Send([]byte("New Book"))
}

func DeleteTodo(c *fiber.Ctx) error {
	return c.Send([]byte("Deletes a book"))
}
