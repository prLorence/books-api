package books

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prLorence/books-api/internal/utils"
)

func SetupRoutes(srv *utils.Server) {
	srv.App.Post("/books", srv.RequireAuth, srv.RequireAdmin, func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return Create(c, srv.DB)
	})

	srv.App.Get("/books", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return All(c, srv.DB)
	})

	srv.App.Get("/books/:id", srv.RequireAuth, func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return GetById(c, srv.DB)
	})

	srv.App.Get("/books/:title", srv.RequireAuth, func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return GetByTitle(c, srv.DB)
	})

	srv.App.Put("/books/:id", srv.RequireAuth, srv.RequireAdmin, func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return Update(c, srv.DB)
	})

	srv.App.Delete("/books/:id", srv.RequireAuth, srv.RequireAdmin, func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return RemoveById(c, srv.DB)
	})
}
