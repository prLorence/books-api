package authors

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prLorence/books-api/internal/db"
)

func GetAuthor(ctx *fiber.Ctx, pg *db.Queries) error {
	return ctx.Send([]byte("Hello World"))
}

func CreateAuthor(ctx *fiber.Ctx, pg *db.Queries) error {
	return ctx.Send(ctx.Body())
}
