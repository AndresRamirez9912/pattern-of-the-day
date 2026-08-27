-- name: CreateAttempt :one
INSERT INTO attempts (
    feedback_id,
    challenge_id,
    status
)
VALUES (?, ?, ?)
RETURNING *;

-- This query fetches the attempts made by a specific user for a specific challenge.
-- It can be used to track the progress of a user on a particular challenge
-- It is obtained by joining the challenges table with the attempts table, filtering by user_id and challenge_id.
-- name: ListAttemptsByUserChallenge :many
SELECT *
FROM attempts
JOIN challenges ON attempts.challenge_id = challenges.id
WHERE challenges.user_id = ? AND challenges.id = ?
ORDER BY attempts.created_at DESC;

-- name: ListAttemptsByChallengeId :many
SELECT *
FROM attempts
WHERE challenge_id = ?
ORDER BY attempts.created_at DESC;

-- name: GetAttemptById :one
SELECT *
FROM attempts
WHERE id = ?;

-- name: UpdateAttempt :one
UPDATE attempts
SET feedback_id = ?,
    status = ?,
    completed_at = CURRENT_TIMESTAMP
WHERE id = ?
RETURNING *;
