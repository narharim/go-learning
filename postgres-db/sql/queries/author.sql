-- queries/001_author.sql

-- name: CreateAuthor :one
INSERT INTO authors (
  name, bio
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateAuthor :one
UPDATE authors
SET name = $2, bio = $3
WHERE id = $1
RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = $1;

-- name: GetAuthor :one
SELECT * FROM authors
WHERE id = $1;

-- name: ListAuthors :many
SELECT * FROM authors
ORDER BY name;

