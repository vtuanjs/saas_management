-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (name, email, avatar, created_at, updated_at)
VALUES ($1, $2, $3, NOW(), NOW())
RETURNING *;