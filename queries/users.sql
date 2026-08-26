-- name: CreateUser :one
INSERT INTO users (
    username,
    email
) 
VALUES (?, ?)
RETURNING *;

-- name: GetUserById :one
SELECT *
FROM users
WHERE id = ?;

-- name: GetUserByUsername :one
SELECT *
FROM users
WHERE username = ?;

-- name: ListUsers :many
SELECT *
FROM users
ORDER BY created_at DESC;

-- name: UpdateUser :one
UPDATE users
SET username = ?, email = ?
WHERE id = ?
RETURNING *;

-- name: ListUserChallenges :many
SELECT *
FROM challenges
WHERE user_id = ?
ORDER BY created_at DESC;
