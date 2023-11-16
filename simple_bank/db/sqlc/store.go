package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(q *Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountId int64 `json:"from_account_id"`
	ToAccountId   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

func (store *Store) TransferEx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var res TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		res.Transfer, err = q.CreatTransfer(ctx, CreatTransferParams{
			FromAccountID: arg.FromAccountId,
			ToAccountID:   arg.ToAccountId,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}
		res.FromEntry, err = q.CreatEntry(ctx, CreatEntryParams{
			AccountID: arg.FromAccountId,
			Amount:    -arg.Amount,	
		})
		if err != nil {
			return err
		}

		res.ToEntry, err = q.CreatEntry(ctx, CreatEntryParams{
			AccountID: arg.ToAccountId,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		if arg.FromAccountId < arg.ToAccountId {
			res.FromAccount, err = q.AddBalance(ctx, AddBalanceParams{
				Amount: -arg.Amount,
				ID:     arg.FromAccountId,
			})
			if err != nil {
				return err
			}
	
			res.ToAccount, err = q.AddBalance(ctx, AddBalanceParams{
				Amount: arg.Amount,
				ID:     arg.ToAccountId,
			})
			if err != nil {
				return err
			}
		} else {
			res.ToAccount, err = q.AddBalance(ctx, AddBalanceParams{
				Amount: arg.Amount,
				ID:     arg.ToAccountId,
			})
			if err != nil {
				return err
			}

			res.FromAccount, err = q.AddBalance(ctx, AddBalanceParams{
				Amount: -arg.Amount,
				ID:     arg.FromAccountId,
			})
			if err != nil {
				return err
			}
		}
		

		return nil
	})

	return res, err
}
