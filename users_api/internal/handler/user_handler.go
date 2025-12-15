package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"

	"users_api/internal/models"
	"users_api/internal/service"
)

type UserHandler struct {
	svc *service.UserService
	val *validator.Validate
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
		val: validator.New(),
	}
}

/* ---------------- CREATE USER ---------------- */

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req models.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}

	if err := h.val.Struct(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	id, err := h.svc.CreateUser(c.Context(), req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{
		"id":   id,
		"name": req.Name,
		"dob":  req.DOB,
	})
}

/* ---------------- GET USER ---------------- */

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id64, err := strconv.ParseInt(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	user, err := h.svc.GetUser(c.Context(), int32(id64))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "user not found"})
	}

	return c.JSON(user)
}

/* ---------------- UPDATE USER ---------------- */

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id64, err := strconv.ParseInt(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	var req models.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}

	if err := h.val.Struct(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.svc.UpdateUser(c.Context(), int32(id64), req); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"id":   id64,
		"name": req.Name,
		"dob":  req.DOB,
	})
}

/* ---------------- DELETE USER ---------------- */

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id64, err := strconv.ParseInt(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	if err := h.svc.DeleteUser(c.Context(), int32(id64)); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}


/* ---------------- LIST ALL USERS ---------------- */

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	users, err := h.svc.ListUsers(
		c.Context(),
		int32(limit),
		int32(offset),
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(users)
}

