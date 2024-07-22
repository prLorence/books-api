package books

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/prLorence/books-api/internal/db"
)

type BookDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	AuthorID    int    `json:"author_id"`
}

func Create(c *fiber.Ctx, pg *db.Queries) error {
	req := new(BookDto)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	book, err := pg.InsertBook(context.TODO(), db.InsertBookParams{
		Title:       req.Title,
		Description: req.Description,
		AuthorID:    int32(req.AuthorID),
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	return c.JSON(book)
}

func All(c *fiber.Ctx, pg *db.Queries) error {
	books, err := pg.SelectBooks(context.TODO())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	return c.JSON(books)
}

func GetById(c *fiber.Ctx, pg *db.Queries) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	book, err := pg.SelectBook(context.TODO(), int32(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	return c.JSON(book)
}

func GetByTitle(c *fiber.Ctx, pg *db.Queries) error {
	title := c.Params("title")
	book, err := pg.SelectByTitle(context.TODO(), title)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	return c.JSON(book)
}

func Update(c *fiber.Ctx, pg *db.Queries) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	req := new(BookDto)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	updBook := pg.UpdateBook(context.TODO(), db.UpdateBookParams{
		ID:          int32(id),
		Title:       req.Title,
		Description: req.Description,
	})

	return c.JSON(updBook)
}

func RemoveById(c *fiber.Ctx, pg *db.Queries) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	err = pg.DeleteBook(context.TODO(), int32(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	return nil
}
