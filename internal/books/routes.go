package books

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prLorence/books-api/internal/db"
)

func SetupRoutes(a *fiber.App, db *db.Queries) {
	a.Post("/books", func(c *fiber.Ctx) error {
		return Create(c, db)
	})

	a.Get("/books", func(c *fiber.Ctx) error {
		return All(c, db)
	})

	a.Get("/books/:id", func(c *fiber.Ctx) error {
		return GetById(c, db)
	})

	a.Get("/books/:title", func(c *fiber.Ctx) error {
		return GetByTitle(c, db)
	})

	a.Put("/books/:id", func(c *fiber.Ctx) error {
		return Update(c, db)
	})

	a.Delete("/books/:id", func(c *fiber.Ctx) error {
		return RemoveById(c, db)
	})
}
