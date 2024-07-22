package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prLorence/books-api/internal/utils"
)

func SetupRoutes(srv *utils.Server) {
	srv.App.Post("/login", func(c *fiber.Ctx) error {
		return Login(c, srv.DB, srv.Session)
	})

	srv.App.Get("/logout", srv.RequireAuth, func(c *fiber.Ctx) error {
		return Logout(c, srv.DB, srv.Session)
	})
}
