-- name: CreateUserChallenge :one
INSERT INTO user_challenges (
    user_id,
    challenge_id,
    status
)
VALUES (?, ?, ?)
RETURNING *;

-- name: GetUserChallengesByUserId :many
SELECT *
FROM user_challenges
WHERE user_id = ?
ORDER BY created_at DESC;

-- name: GetUserChallengesByChallengeId :many
SELECT *
FROM user_challenges
WHERE challenge_id = ?
ORDER BY created_at DESC;

-- name: GetUserChallengeById :one
SELECT *
FROM user_challenges
WHERE id = ?;

-- name: UpdateUserChallengeStatus :one
UPDATE user_challenges
SET status = ?
WHERE id = ?
RETURNING *;
