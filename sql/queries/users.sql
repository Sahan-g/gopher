-- name: CreateUser :one 
INSERT INTO users (username, email,api_key) VALUES ($1, $2, 
encode(sha256(random()::text::bytea), 'hex')) RETURNING *;

-- name: UserByEmail :one
SELECT id FROM users WHERE email = $1;

-- name: UserByApiKey :one
SELECT * FROM users WHERE api_key = $1; 