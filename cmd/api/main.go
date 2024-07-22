package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/storage/postgres/v3"
	"github.com/prLorence/books-api/config"
	"github.com/prLorence/books-api/internal/db"
	"github.com/prLorence/books-api/internal/utils"
)

func main() {
	config := config.NewConfig()
	conn, err := db.ConnectDB(config.DB_CONN)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v\n", err)
	}

	srv := &utils.Server{
		App: fiber.New(),
		DB:  db.New(conn),
		// gofiber's postgresql driver for session storage
		Session: postgres.New(postgres.Config{
			ConnectionURI: config.DB_CONN,
		}),
	}
	srv.App.Use(cors.New())

	log.Println("Seeding users")
	utils.SeedUsers(srv)
	Routes(srv)

	fmt.Fprintf(os.Stdout, "Listening on host port %s", config.HOST_PORT)
	err = srv.App.Listen(fmt.Sprintf(":%s", config.APP_PORT))
	log.Fatal(err)
}
