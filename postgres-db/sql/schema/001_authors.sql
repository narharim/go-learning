-- +goose Up
CREATE TABLE authors (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  bio TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE authors;
