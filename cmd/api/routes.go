package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/prLorence/books-api/internal/authors"
	"github.com/prLorence/books-api/internal/db"
)

func Routes(app *fiber.App, db *db.Queries) http.Handler {
	app.Post("/authors", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return authors.CreateAuthor(c, db)
	})

	app.Get("/authors", func(c *fiber.Ctx) error {
		return authors.GetAuthor(c, db)
	})
	return nil
}
