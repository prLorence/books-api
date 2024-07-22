package users

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/storage/postgres/v3"
	"github.com/prLorence/books-api/internal/db"
	"github.com/prLorence/books-api/internal/utils"
)

// login sets the session of user id
// logout removes the user id from the session
type UserDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx, pg *db.Queries, store *postgres.Storage) error {
	req := UserDto{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	id, err := pg.SelectUser(context.TODO(), db.SelectUserParams{
		UserName:     req.Username,
		PasswordHash: req.Password,
	})
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	store.Set("userId", utils.Int32ToBytes(id), 12*time.Hour)

	return c.JSON(fiber.Map{
		"userId": id,
	})
}

func Logout(c *fiber.Ctx, pg *db.Queries, store *postgres.Storage) error {
	return store.Delete("userId")
}
