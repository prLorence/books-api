package books

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prLorence/books-api/internal/db"
	"github.com/prLorence/books-api/internal/utils"
)

func SetupRoutes(a *fiber.App, db *db.Queries) {
	a.Post("/books", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return Create(c, db)
	}).Use(adaptor.HTTPMiddleware(utils.RequireAdmin))

	a.Get("/books", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return All(c, db)
	})

	a.Get("/books/:id", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return GetById(c, db)
	})

	a.Get("/books/:title", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return GetByTitle(c, db)
	})

	a.Put("/books/:id", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return Update(c, db)
	}).Use(adaptor.HTTPMiddleware(utils.RequireAdmin))

	a.Delete("/books/:id", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return RemoveById(c, db)
	}).Use(adaptor.HTTPMiddleware(utils.RequireAdmin))
}
