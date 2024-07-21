-- name: 
-- name: GetAuthors :many
SELECT * FROM authors
ORDER BY id;

-- name: GetAuthor :one
SELECT * FROM authors
WHERE id=$1 LIMIT 1;

-- name: InsertAuthor :one
INSERT INTO authors (name)
VALUES ($1)
RETURNING *;

-- name: UpdateAuthor :exec
UPDATE authors 
SET name=$2
WHERE id=$1;

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id=$1;

-- name: GetBooks :many
SELECT * FROM books
ORDER by title;

-- name: GetBook :one
SELECT * FROM books
WHERE id=$1 LIMIT 1;

-- name: InsertBook :one
INSERT INTO books (author_id, title, description)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateBook :one
UPDATE books
SET title=$2, description=$3
WHERE id=$1;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id=$1;

