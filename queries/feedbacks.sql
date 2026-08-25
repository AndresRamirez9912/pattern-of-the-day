-- name: CreateFeedback :one
INSERT INTO feedbacks (
    score,
    rating,
    summary,
    suggestions
)
VALUES (?, ?, ?, ?)
RETURNING *;

-- name: GetFeedbackById :one
SELECT *
FROM feedbacks
WHERE id = ?;
