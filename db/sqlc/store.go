package db

import (
	"context"
	"database/sql"
	"fmt"
)

// ? Store provides all functions to execute db aueries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// ? ExecTx executes a function within a database transaction
// ? Create a new Queries object with that transaction and calls callback function with the created queries
// ? and finally commit or rollback the transaction based on the error that returnes from that functoin
// this functin is unexported
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	//starts new transactoin with calling TxOptions
	// this option allows us  to set custom isolaiton level of this transation
	// and if we don't set it explicitly then deafual isolation level will be used
	tx, err := store.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// ? TransferTxParams contains all the input parameters of the transfer transaction
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// ? TransferTxResult is the result of the transfer transaction
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// ? TransferTx perform a money transfer form one account to the other.
// ? It creates a tansfer record, add account enteries, and update account's balance witin a single database tansactions
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	// this is callback function
	err := store.execTx(ctx, func(q *Queries) error {
		// queries
		var err error
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}
		// todo: update account balance
		return nil
	})

	return result, err
}
