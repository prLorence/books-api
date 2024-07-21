package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prLorence/books-api/internal/authors"
	"github.com/prLorence/books-api/internal/db"
)

func Routes(app *fiber.App, db *db.Queries) {
	authors.SetupRoutes(app, db)
}
