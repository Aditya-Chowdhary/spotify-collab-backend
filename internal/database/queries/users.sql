-- name: CreateUser :one
INSERT INTO users(name, email, password_hash, activated)
VALUES ($1, $2, $3, $4)
RETURNING user_uuid, id, created_at, version;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;

-- name: GetUserByUUID :one
SELECT *
FROM users
WHERE user_uuid = $1;

-- name: UpdateUser :one
UPDATE users
SET name = $1, email = $2, password_hash = $3, activated = $4, version = version + 1
WHERE id=$5 AND version = $6
RETURNING *;


-- name: DeleteUser :exec
DELETE FROM users
WHERE user_uuid = $1;