package authors

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prLorence/books-api/internal/utils"
)

func SetupRoutes(srv *utils.Server) {
	srv.App.Post("/authors", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return Create(c, srv.DB)
	}).Use(adaptor.HTTPMiddleware(srv.RequireAdmin))

	srv.App.Get("/authors", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return All(c, srv.DB)
	})

	srv.App.Get("/authors/:id", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return GetById(c, srv.DB)
	})

	srv.App.Put("/authors/:id", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return Update(c, srv.DB)
	}).Use(adaptor.HTTPMiddleware(srv.RequireAdmin))

	srv.App.Delete("/authors/:id", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return RemoveById(c, srv.DB)
	}).Use(adaptor.HTTPMiddleware(srv.RequireAdmin))
}
