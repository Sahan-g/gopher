-- name: CreateUser :one 
INSERT INTO users (username, email) VALUES ($1, $2) RETURNING id, username, email, created_at;

-- name: UserByEmail :one
SELECT id FROM users WHERE email = $1 LIMIT 1;