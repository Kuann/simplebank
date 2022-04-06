// Code generated by sqlc. DO NOT EDIT.
// source: account.sql

package db

import (
	"context"
)

const createAccount = `-- name: CreateAccount :one
insert into account (username, password) values ($1, $2) returning id, username, password
`

type CreateAccountParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount, arg.Username, arg.Password)
	var i Account
	err := row.Scan(&i.ID, &i.Username, &i.Password)
	return i, err
}

const getAccount = `-- name: GetAccount :one
SELECT id, username, password FROM account
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, id int32) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccount, id)
	var i Account
	err := row.Scan(&i.ID, &i.Username, &i.Password)
	return i, err
}

const listAccounts = `-- name: ListAccounts :many
SELECT id, username, password FROM account
`

func (q *Queries) ListAccounts(ctx context.Context) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, listAccounts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Account
	for rows.Next() {
		var i Account
		if err := rows.Scan(&i.ID, &i.Username, &i.Password); err != nil {
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

const listOrderedAccounts = `-- name: ListOrderedAccounts :many
SELECT id, username, password FROM account ORDER BY username OFFSET $1 LIMIT $2
`

type ListOrderedAccountsParams struct {
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

func (q *Queries) ListOrderedAccounts(ctx context.Context, arg ListOrderedAccountsParams) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, listOrderedAccounts, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Account
	for rows.Next() {
		var i Account
		if err := rows.Scan(&i.ID, &i.Username, &i.Password); err != nil {
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