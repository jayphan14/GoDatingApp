-- name: CreateUser :one
INSERT INTO users (
  username, email, password, gender, university, picture, bio, bio_pictures 
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;


-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
set username = $2,
 email = $3, 
 password = $4, 
 gender = $5, 
 university = $6, 
 picture = $7, 
 bio = $8, 
 bio_pictures = $9  

WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;