-- name: CreateChallenge :one
INSERT INTO challenges (
    name,
    description,
    difficulty,
    type,
    target_pattern,
    status,
    created_at,
    updated_at
)
VALUES (?, ?, ?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: GetChallengeById :one
SELECT *
FROM challenges
WHERE id = ?;

-- name: ListChallenges :many
SELECT *
FROM challenges
ORDER BY created_at DESC;
