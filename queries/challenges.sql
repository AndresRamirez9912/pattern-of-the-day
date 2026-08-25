-- name: CreateChallenge :one
INSERT INTO challenges (
    name,
    description,
    difficulty,
    type,
    target_pattern
)
VALUES (?, ?, ?, ?, ?)
RETURNING *;

-- name: GetChallengeById :one
SELECT *
FROM challenges
WHERE id = ?;

-- name: ListChallenges :many
SELECT *
FROM challenges
ORDER BY created_at DESC;

-- name: UpdateChallenge :one
UPDATE challenges
SET name = ?,
    description = ?,
    difficulty = ?,
    type = ?,
    target_pattern = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?
RETURNING *;
