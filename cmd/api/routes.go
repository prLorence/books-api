package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/prLorence/books-api/internal/authors"
	"github.com/prLorence/books-api/internal/books"
	"github.com/prLorence/books-api/internal/db"
	"github.com/prLorence/books-api/internal/users"
)

func Routes(app *fiber.App, db *db.Queries, store *session.Store) {
	authors.SetupRoutes(app, db)
	books.SetupRoutes(app, db)
	users.SetupRoutes(app, db, store)
}
