package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/prLorence/books-api/config"
	"github.com/prLorence/books-api/internal/db"
)

// middleware, for authorization
// decoder from request to application struct
// db connection string comes from os.env, which come from config struct, might need an if statement for dev purposes
// session manager
// a seeder

type Application struct {
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	config := config.NewConfig()
	conn, err := db.ConnectDB(config.DB_CONN)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v\n", err)
	}

	db := db.New(conn)

	Routes(app, db)

	fmt.Fprintf(os.Stdout, "Listening on host port %s", config.HOST_PORT)
	err = app.Listen(fmt.Sprintf(":%s", config.APP_PORT))
	log.Fatal(err)
}
