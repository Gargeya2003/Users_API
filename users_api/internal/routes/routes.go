package routes

import (
	"github.com/gofiber/fiber/v2"
	"users_api/internal/handler"
)

func Register(app *fiber.App, h *handler.UserHandler) {
	users := app.Group("/users")

	users.Post("", h.CreateUser)       // POST /users
	users.Get("/:id", h.GetUser)       // GET /users/:id
	users.Get("", h.ListUsers)         // GET /users
	users.Put("/:id", h.UpdateUser)    // PUT /users/:id
	users.Delete("/:id", h.DeleteUser) // DELETE /users/:id
}
