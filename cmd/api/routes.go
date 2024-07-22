package main

import (
	"github.com/prLorence/books-api/internal/authors"
	"github.com/prLorence/books-api/internal/books"
	"github.com/prLorence/books-api/internal/users"
	"github.com/prLorence/books-api/internal/utils"
)

// func Routes(app *fiber.App, db *db.Queries, store *session.Store) {
// 	authors.SetupRoutes(app, db)
// 	books.SetupRoutes(app, db)
// 	users.SetupRoutes(app, db, store)
// }

func Routes(srv *utils.Server) {
	authors.SetupRoutes(srv.App, srv.DB)
	books.SetupRoutes(srv.App, srv.DB)
	users.SetupRoutes(srv.App, srv.DB, srv.Session)
}
