// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: transfers.sql

package db

import (
	"context"
)

const creatTransfer = `-- name: CreatTransfer :one
INSERT INTO transfers (
    from_account_id,
    to_account_id,
    amount
)
VALUES (
    $1, $2, $3 
)
RETURNING id, from_account_id, to_account_id, amount, created_at
`

type CreatTransferParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

func (q *Queries) CreatTransfer(ctx context.Context, arg CreatTransferParams) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, creatTransfer, arg.FromAccountID, arg.ToAccountID, arg.Amount)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const deleteTransfer = `-- name: DeleteTransfer :exec
DELETE from transfers
WHERE id = $1
`

func (q *Queries) DeleteTransfer(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTransfer, id)
	return err
}

const getAllTransfers = `-- name: GetAllTransfers :many
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
ORDER BY id
LIMIT $1
OFFSET $2
`

type GetAllTransfersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetAllTransfers(ctx context.Context, arg GetAllTransfersParams) ([]Transfer, error) {
	rows, err := q.db.QueryContext(ctx, getAllTransfers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transfer
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTransfer = `-- name: GetTransfer :one
SELECT  id, from_account_id, to_account_id, amount, created_at
FROM transfers
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetTransfer(ctx context.Context, id int64) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, getTransfer, id)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getTransferForUpdate = `-- name: GetTransferForUpdate :one
SELECT  id, from_account_id, to_account_id, amount, created_at
FROM transfers
WHERE id = $1
LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetTransferForUpdate(ctx context.Context, id int64) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, getTransferForUpdate, id)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const updateTransfer = `-- name: UpdateTransfer :one
UPDATE transfers
    set from_account_id = $2,
    to_account_id = $3,
    amount = $4
WHERE id = $1
RETURNING id, from_account_id, to_account_id, amount, created_at
`

type UpdateTransferParams struct {
	ID            int64 `json:"id"`
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

func (q *Queries) UpdateTransfer(ctx context.Context, arg UpdateTransferParams) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, updateTransfer,
		arg.ID,
		arg.FromAccountID,
		arg.ToAccountID,
		arg.Amount,
	)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}