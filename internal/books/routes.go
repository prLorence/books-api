package books

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prLorence/books-api/internal/utils"
)

func SetupRoutes(srv *utils.Server) {
	srv.App.Post("/books", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return Create(c, srv.DB)
	}).Use(adaptor.HTTPMiddleware(srv.RequireAdmin))

	srv.App.Get("/books", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return All(c, srv.DB)
	})

	srv.App.Get("/books/:id", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return GetById(c, srv.DB)
	})

	srv.App.Get("/books/:title", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return GetByTitle(c, srv.DB)
	})

	srv.App.Put("/books/:id", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return Update(c, srv.DB)
	}).Use(adaptor.HTTPMiddleware(srv.RequireAdmin))

	srv.App.Delete("/books/:id", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return RemoveById(c, srv.DB)
	}).Use(adaptor.HTTPMiddleware(srv.RequireAdmin))
}
