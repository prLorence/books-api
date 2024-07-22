package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/prLorence/books-api/config"
	"github.com/prLorence/books-api/internal/db"
	"github.com/prLorence/books-api/internal/utils"
)

// middleware, for authorization
// decoder from request to application struct
// db connection string comes from os.env, which come from config struct, might need an if statement for dev purposes
// session manager
// a seeder

func main() {
	config := config.NewConfig()
	conn, err := db.ConnectDB(config.DB_CONN)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v\n", err)
	}

	srv := &utils.Server{
		App:     fiber.New(),
		DB:      db.New(conn),
		Session: session.New(),
	}
	srv.App.Use(cors.New())
	srv.App.Use(adaptor.HTTPMiddleware(srv.RequireAuth))

	// Routes(srv.App, srv.DB, srv.Session)
	Routes(srv)

	fmt.Fprintf(os.Stdout, "Listening on host port %s", config.HOST_PORT)
	err = srv.App.Listen(fmt.Sprintf(":%s", config.APP_PORT))
	log.Fatal(err)
}
