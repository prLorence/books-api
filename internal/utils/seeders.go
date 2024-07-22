package utils

import (
	"context"

	"github.com/prLorence/books-api/internal/db"
)

func SeedUsers(srv *Server) {
	adminId, _ := srv.DB.InsertUser(context.TODO(), db.InsertUserParams{
		UserName:     "admin",
		PasswordHash: "admin",
	})

	userId, _ := srv.DB.InsertUser(context.TODO(), db.InsertUserParams{
		UserName:     "user",
		PasswordHash: "user",
	})

	adminRoleId, _ := srv.DB.InsertRole(context.TODO(), "admin")
	userRoleId, _ := srv.DB.InsertRole(context.TODO(), "user")

	srv.DB.InsertUserRole(context.TODO(), db.InsertUserRoleParams{
		UserID: adminId,
		RoleID: adminRoleId,
	})

	srv.DB.InsertUserRole(context.TODO(), db.InsertUserRoleParams{
		UserID: userId,
		RoleID: userRoleId,
	})
}
