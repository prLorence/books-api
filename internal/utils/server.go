package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/prLorence/books-api/internal/db"
)

type Server struct {
	App     *fiber.App
	DB      *db.Queries
	Session *session.Store

}
