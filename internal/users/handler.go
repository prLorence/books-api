package users

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/prLorence/books-api/internal/db"
)

// login sets the session of user id
// logout removes the user id from the session
type UserDto struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx, pg *db.Queries, store *session.Store) error {
	req := UserDto{}
	if err := c.BodyParser(req); err != nil {
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

	// sess, err := store.Get(c)

	return c.JSON(id)
}
