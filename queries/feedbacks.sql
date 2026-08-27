-- name: CreateFeedback :one
INSERT INTO feedbacks (
    score,
    summary,
    suggestions
)
VALUES ( ?, ?, ?)
RETURNING *;

-- name: GetFeedbackById :one
SELECT *
FROM feedbacks
WHERE id = ?;
