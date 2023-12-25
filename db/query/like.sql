-- name: CreateUser :one
INSERT INTO likes (
    "senderID", "receiverID"
) VALUES (
  $1, $2
)
RETURNING *;