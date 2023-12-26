-- name: CreateMatch :one
INSERT INTO matches (
  user1id, user2id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetMatchByMatchID :one
SELECT * FROM matches
WHERE id = $1 LIMIT 1;

-- name: GetMatchByUserId :many
SELECT * FROM matches
WHERE (user1id = $1 AND user2id = $2) OR (user1id = $2 AND user2id = $1);

-- name: ListMatches :many
SELECT * FROM matches
WHERE user1id = $1 OR user2id = $2;

-- name: DeleteMatch :exec
DELETE FROM matches
WHERE (user1id = $1 AND user2id = $2) OR (user1id = $2 AND user2id = $1);
