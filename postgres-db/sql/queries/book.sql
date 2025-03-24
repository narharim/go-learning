-- queries/002_book.sql

-- name: CreateBook :one
INSERT INTO books (
  author_id, title, description, published_year
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateBook :one
UPDATE books
SET title = $2, description = $3, published_year = $4, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1;

-- name: GetBook :one
SELECT * FROM books
WHERE id = $1;

-- name: ListBooksByAuthor :many
SELECT b.* FROM books b
JOIN authors a ON b.author_id = a.id
WHERE a.id = $1
ORDER BY b.title;
