package utils

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) RequireAuth(c *fiber.Ctx) error {
	store := s.Session

	id, _ := store.Get("userId")
	if BytesToInt32(id) == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "You are not logged in",
		})
	}

	return c.Next()
}

func (s *Server) RequireAdmin(c *fiber.Ctx) error {
	store := s.Session

	idInBytes, err := store.Get("userId")
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "You are not logged in",
		})
	}

	isAdmin, err := s.DB.IsAdmin(context.TODO(), BytesToInt32(idInBytes))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	if !isAdmin {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "You don't have permissions to access this resource",
		})
	}

	return c.Next()
}
