package authors

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prLorence/books-api/internal/db"
)

func SetupRoutes(app *fiber.App, db *db.Queries) {
	app.Post("/authors", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return Create(c, db)
	})

	app.Get("/authors", func(c *fiber.Ctx) error {
		return All(c, db)
	})

	app.Get("/authors/:id", func(c *fiber.Ctx) error {
		return GetById(c, db)
	})

	app.Put("/authors/:id", func(c *fiber.Ctx) error {
		return Update(c, db)
	})

	app.Delete("/authors/:id", func(c *fiber.Ctx) error {
		return RemoveById(c, db)
	})
}
