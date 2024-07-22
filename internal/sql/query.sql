-- name: SelectAuthors :many
SELECT * FROM authors
ORDER BY id;

-- name: SelectAuthor :one
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

-- name: AuthorExists :exec
SELECT EXISTS(
SELECT true FROM authors WHERE name=$1);

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id=$1;

-- name: SelectBooks :many
SELECT b.title, b.description,
    a.name AS author_name
FROM Books b
JOIN Authors a ON b.author_id = a.id;

-- name: SelectBook :one
SELECT b.title, b.description,
    a.name AS author_name
FROM Books b
JOIN Authors a ON b.author_id = a.id
WHERE b.id=$1 LIMIT 1;

-- name: SelectByTitle :one
SELECT * FROM books
WHERE title=$1 LIMIT 1;

-- name: InsertBook :one
INSERT INTO books (author_id, title, description)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateBook :exec
UPDATE books
SET title=$2, description=$3
WHERE id=$1;

-- name: BookExists :exec
SELECT EXISTS(
SELECT true FROM books WHERE title=$1);

-- name: DeleteBook :exec
DELETE FROM books
WHERE id=$1;

-- name: InsertRole :one
INSERT INTO roles (name)
VALUES ($1)
ON CONFLICT DO NOTHING
RETURNING id;

-- name: InsertUser :one
INSERT INTO users (user_name, password_hash)
VALUES ($1, $2)
ON CONFLICT DO NOTHING
RETURNING id;

-- name: SelectUser :one
SELECT id from users 
WHERE user_name=$1 AND password_hash=$2;

-- name: InsertUserRole :exec
INSERT INTO userroles (user_id, role_id)
VALUES ($1, $2)
ON CONFLICT DO NOTHING;

-- name: UserExists :one
SELECT EXISTS(
SELECT true FROM users WHERE user_name=$1);

-- name: GetUserHash :one
SELECT id, password_hash FROM users
WHERE user_name=$1;

-- name: IsAdmin :one
SELECT EXISTS (
    SELECT 1
    FROM UserRoles ur
    JOIN Roles r ON ur.role_id = r.id
    WHERE ur.user_id = $1
    AND r.name = 'admin'
) AS has_admin_role;
