-- name: CreateLike :one
INSERT INTO likes (
  sender_id, receiver_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetLike :one
SELECT * FROM likes
WHERE id = $1 LIMIT 1;

-- name: ListLikesReceived :many
SELECT * FROM likes
WHERE receiver_id = $1;

-- name: ListLikesSent :many
SELECT * FROM likes
WHERE sender_id = $1;

-- name: DeleteLike :exec
DELETE FROM likes
WHERE sender_id = $1 AND receiver_id = $2;

