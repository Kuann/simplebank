package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("rb error")
		}
		return err
	}
	return tx.Commit()
}

func (store *Store) CreateAccountTx(ctx context.Context, arg CreateAccountParams) (Account, error) {
	var account Account
	return account, store.execTx(ctx, func(q *Queries) error {
		var err error
		account, err = q.CreateAccount(ctx, arg)
		return err
	})
}
