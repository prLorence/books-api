package authors

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/prLorence/books-api/internal/db"
)

type AuthorDto struct {
	Name string `json:"name"`
}

func Create(c *fiber.Ctx, pg *db.Queries) error {
	req := new(AuthorDto)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	author, err := pg.InsertAuthor(context.TODO(), req.Name)

	if err != nil {
		return err
	}

	return c.JSON(author)
}

func All(c *fiber.Ctx, pg *db.Queries) error {
	authors, err := pg.SelectAuthors(context.TODO())

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	return c.JSON(authors)
}

func GetById(c *fiber.Ctx, pg *db.Queries) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	author, err := pg.SelectAuthor(context.TODO(), int32(id))

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	return c.JSON(author)
}

func Update(c *fiber.Ctx, pg *db.Queries) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	req := new(AuthorDto)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	updParams := db.UpdateAuthorParams{
		ID:   int32(id),
		Name: req.Name,
	}

	updAuthor := pg.UpdateAuthor(context.TODO(), updParams)

	return c.JSON(updAuthor)
}

func RemoveById(c *fiber.Ctx, pg *db.Queries) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	err = pg.DeleteAuthor(context.TODO(), int32(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	return nil
}
