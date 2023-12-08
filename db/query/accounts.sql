-- name: GetAccount :one
SELECT  *
FROM accounts
WHERE id = $1
LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT  *
FROM accounts
WHERE id = $1
LIMIT 1
FOR NO KEY UPDATE;

-- name: AddBalance :one
UPDATE accounts
SET balance = balance + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: CreatAccount :one
INSERT INTO accounts (
    owner,
    balance,
    currency 
)
VALUES (
    $1, $2, $3 
)
RETURNING *;

-- name: GetAllAccountsByOwner :many
SELECT * FROM accounts
WHERE owner = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: GetAllAccounts :many
SELECT * FROM accounts
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateAccount :one
UPDATE accounts
    set owner = $2,
    balance = $3, 
    currency = $4
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE from accounts
WHERE id = $1;