package authors

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prLorence/books-api/internal/db"
	"github.com/prLorence/books-api/internal/utils"
)

func SetupRoutes(app *fiber.App, db *db.Queries) {
	app.Post("/authors", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return Create(c, db)
	}).Use(adaptor.HTTPMiddleware(utils.RequireAdmin))

	app.Get("/authors", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return All(c, db)
	})

	app.Get("/authors/:id", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return GetById(c, db)
	})

	app.Put("/authors/:id", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return Update(c, db)
	}).Use(adaptor.HTTPMiddleware(utils.RequireAdmin))

	app.Delete("/authors/:id", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return RemoveById(c, db)
	}).Use(adaptor.HTTPMiddleware(utils.RequireAdmin))
}
