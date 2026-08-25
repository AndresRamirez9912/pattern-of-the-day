-- name: CreateAttempt :one
INSERT INTO attempts (
    feedback_id,
    user_challenge_id,
    status
)
VALUES (?, ?, ?)
RETURNING *;

-- This query fetches the attemps made by a specific user for a specific challenge.
-- It can be used to track the progress of a user on a particular challenge
-- It is obtained joinning the table user_challenges with the attempts table, filtering by user_id and challenge_id.
-- name: GetUserAttemptsByChallengeId :many
SELECT *
FROM attempts
JOIN user_challenges ON attempts.user_challenge_id = user_challenges.id
WHERE user_challenges.user_id = ? AND user_challenges.challenge_id = ?
ORDER BY attempts.created_at DESC;


