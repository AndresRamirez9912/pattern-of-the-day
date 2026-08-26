-- name: CreateClue :one
INSERT INTO clues (
    challenge_id,
    description,
    sequence_order
)
VALUES (?, ?, ?)
RETURNING *;

-- name: GetCluesByChallengeId :many
SELECT *
FROM clues
WHERE challenge_id = ?
ORDER BY sequence_order ASC;

-- name: ListCluesByChallengeId :many
SELECT *
FROM clues
WHERE challenge_id = ?
ORDER BY sequence_order ASC;
