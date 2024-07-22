package authors

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prLorence/books-api/internal/utils"
)

func SetupRoutes(srv *utils.Server) {
	srv.App.Post("/authors", srv.RequireAuth, srv.RequireAdmin, func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return Create(c, srv.DB)
	})

	srv.App.Get("/authors", srv.RequireAuth, func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return All(c, srv.DB)
	})

	srv.App.Get("/authors/:id", srv.RequireAuth, func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return GetById(c, srv.DB)
	})

	srv.App.Put("/authors/:id", srv.RequireAuth, srv.RequireAdmin, func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return Update(c, srv.DB)
	})

	srv.App.Delete("/authors/:id", srv.RequireAuth, srv.RequireAdmin, func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return RemoveById(c, srv.DB)
	})
}
