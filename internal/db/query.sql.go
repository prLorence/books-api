// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package db

import (
	"context"
)

const authorExists = `-- name: AuthorExists :exec
SELECT EXISTS(
SELECT true FROM authors WHERE name=$1)
`

func (q *Queries) AuthorExists(ctx context.Context, name string) error {
	_, err := q.db.Exec(ctx, authorExists, name)
	return err
}

const bookExists = `-- name: BookExists :exec
SELECT EXISTS(
SELECT true FROM books WHERE title=$1)
`

func (q *Queries) BookExists(ctx context.Context, title string) error {
	_, err := q.db.Exec(ctx, bookExists, title)
	return err
}

const deleteAuthor = `-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id=$1
`

func (q *Queries) DeleteAuthor(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteAuthor, id)
	return err
}

const deleteBook = `-- name: DeleteBook :exec
DELETE FROM books
WHERE id=$1
`

func (q *Queries) DeleteBook(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteBook, id)
	return err
}

const getUserHash = `-- name: GetUserHash :one
SELECT id, password_hash FROM users
WHERE user_name=$1
`

type GetUserHashRow struct {
	ID           int32
	PasswordHash string
}

func (q *Queries) GetUserHash(ctx context.Context, userName string) (GetUserHashRow, error) {
	row := q.db.QueryRow(ctx, getUserHash, userName)
	var i GetUserHashRow
	err := row.Scan(&i.ID, &i.PasswordHash)
	return i, err
}

const insertAuthor = `-- name: InsertAuthor :one
INSERT INTO authors (name)
VALUES ($1)
RETURNING id, name
`

func (q *Queries) InsertAuthor(ctx context.Context, name string) (Author, error) {
	row := q.db.QueryRow(ctx, insertAuthor, name)
	var i Author
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const insertBook = `-- name: InsertBook :one
INSERT INTO books (author_id, title, description)
VALUES ($1, $2, $3)
RETURNING id, author_id, title, description
`

type InsertBookParams struct {
	AuthorID    int32
	Title       string
	Description string
}

func (q *Queries) InsertBook(ctx context.Context, arg InsertBookParams) (Book, error) {
	row := q.db.QueryRow(ctx, insertBook, arg.AuthorID, arg.Title, arg.Description)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.AuthorID,
		&i.Title,
		&i.Description,
	)
	return i, err
}

const isAdmin = `-- name: IsAdmin :one
SELECT EXISTS (
    SELECT 1
    FROM UserRoles ur
    JOIN Roles r ON ur.role_id = r.id
    WHERE ur.user_id = $1
    AND r.name = 'admin'
) AS has_admin_role
`

func (q *Queries) IsAdmin(ctx context.Context, userID int32) (bool, error) {
	row := q.db.QueryRow(ctx, isAdmin, userID)
	var has_admin_role bool
	err := row.Scan(&has_admin_role)
	return has_admin_role, err
}

const seedRoles = `-- name: SeedRoles :exec
INSERT INTO roles (name)
VALUES ('admin'), ('user')
ON CONFLICT DO NOTHING
`

func (q *Queries) SeedRoles(ctx context.Context) error {
	_, err := q.db.Exec(ctx, seedRoles)
	return err
}

const seedUserRoles = `-- name: SeedUserRoles :exec
INSERT INTO userroles (user_id, role_id)
VALUES ($1, $2)
ON CONFLICT DO NOTHING
`

type SeedUserRolesParams struct {
	UserID int32
	RoleID int32
}

func (q *Queries) SeedUserRoles(ctx context.Context, arg SeedUserRolesParams) error {
	_, err := q.db.Exec(ctx, seedUserRoles, arg.UserID, arg.RoleID)
	return err
}

const seedUsers = `-- name: SeedUsers :exec
INSERT INTO users (user_name, password_hash)
VALUES ($1, $2)
ON CONFLICT DO NOTHING
`

type SeedUsersParams struct {
	UserName     string
	PasswordHash string
}

func (q *Queries) SeedUsers(ctx context.Context, arg SeedUsersParams) error {
	_, err := q.db.Exec(ctx, seedUsers, arg.UserName, arg.PasswordHash)
	return err
}

const selectAuthor = `-- name: SelectAuthor :one
SELECT id, name FROM authors
WHERE id=$1 LIMIT 1
`

func (q *Queries) SelectAuthor(ctx context.Context, id int32) (Author, error) {
	row := q.db.QueryRow(ctx, selectAuthor, id)
	var i Author
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const selectAuthors = `-- name: SelectAuthors :many
SELECT id, name FROM authors
ORDER BY id
`

func (q *Queries) SelectAuthors(ctx context.Context) ([]Author, error) {
	rows, err := q.db.Query(ctx, selectAuthors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectBook = `-- name: SelectBook :one
SELECT id, author_id, title, description FROM books
WHERE id=$1 LIMIT 1
`

func (q *Queries) SelectBook(ctx context.Context, id int32) (Book, error) {
	row := q.db.QueryRow(ctx, selectBook, id)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.AuthorID,
		&i.Title,
		&i.Description,
	)
	return i, err
}

const selectBooks = `-- name: SelectBooks :many
SELECT id, author_id, title, description FROM books
ORDER by title
`

func (q *Queries) SelectBooks(ctx context.Context) ([]Book, error) {
	rows, err := q.db.Query(ctx, selectBooks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Book
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.ID,
			&i.AuthorID,
			&i.Title,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectByTitle = `-- name: SelectByTitle :one
SELECT id, author_id, title, description FROM books
WHERE title=$1 LIMIT 1
`

func (q *Queries) SelectByTitle(ctx context.Context, title string) (Book, error) {
	row := q.db.QueryRow(ctx, selectByTitle, title)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.AuthorID,
		&i.Title,
		&i.Description,
	)
	return i, err
}

const selectUser = `-- name: SelectUser :one
SELECT id from users 
WHERE user_name=$1 AND password_hash=$2
`

type SelectUserParams struct {
	UserName     string
	PasswordHash string
}

func (q *Queries) SelectUser(ctx context.Context, arg SelectUserParams) (int32, error) {
	row := q.db.QueryRow(ctx, selectUser, arg.UserName, arg.PasswordHash)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const updateAuthor = `-- name: UpdateAuthor :exec
UPDATE authors 
SET name=$2
WHERE id=$1
`

type UpdateAuthorParams struct {
	ID   int32
	Name string
}

func (q *Queries) UpdateAuthor(ctx context.Context, arg UpdateAuthorParams) error {
	_, err := q.db.Exec(ctx, updateAuthor, arg.ID, arg.Name)
	return err
}

const updateBook = `-- name: UpdateBook :exec
UPDATE books
SET title=$2, description=$3
WHERE id=$1
`

type UpdateBookParams struct {
	ID          int32
	Title       string
	Description string
}

func (q *Queries) UpdateBook(ctx context.Context, arg UpdateBookParams) error {
	_, err := q.db.Exec(ctx, updateBook, arg.ID, arg.Title, arg.Description)
	return err
}

const userExists = `-- name: UserExists :one
SELECT EXISTS(
SELECT true FROM users WHERE user_name=$1)
`

func (q *Queries) UserExists(ctx context.Context, userName string) (bool, error) {
	row := q.db.QueryRow(ctx, userExists, userName)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}
