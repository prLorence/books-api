package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/storage/postgres/v3"
	"github.com/prLorence/books-api/internal/db"
)

type Server struct {
	App     *fiber.App
	DB      *db.Queries
	Session *postgres.Storage
}
