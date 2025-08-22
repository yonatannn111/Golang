-- +goose UP
CREATE TABLE users (
    id uuid primary key,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL
);
-- +goose DOWN
DROP TABLE users;