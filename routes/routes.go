package routes

import (
	"go_todo/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// User routes
	api.Post("/users", controllers.CreateUser)
	api.Get("/users", controllers.GetUsers)
	api.Get("/users/:id", controllers.GetUser)
	api.Put("/users/:id", controllers.UpdateUser)
	api.Delete("/users/:id", controllers.DeleteUser)

	// Todo routes
	api.Post("/todos", controllers.CreateTodo)
	api.Get("/todos", controllers.GetTodos)
	api.Get("/todos/:id", controllers.GetTodo)
	api.Put("/todos/:id", controllers.UpdateTodo)
	api.Delete("/todos/:id", controllers.DeleteTodo)
	api.Get("/todos/user/:userId", controllers.GetTodosByUser)
}
