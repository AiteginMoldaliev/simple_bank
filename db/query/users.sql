-- name: GetUser :one
SELECT  *
FROM users
WHERE username = $1
LIMIT 1;

-- name: GetUserForUpdate :one
SELECT  *
FROM users
WHERE username = $1
LIMIT 1
FOR NO KEY UPDATE;

-- name: CreatUser :one
INSERT INTO users (
    username,
	hashed_password,
	full_name,
	email
)
VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: GetAllUsers :many
SELECT * FROM users
ORDER BY username
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
    set username = $2,
    hashed_password = $3, 
    full_name = $4,
    email = $5
WHERE username = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE from users
WHERE username = $1;