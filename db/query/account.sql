-- name: CreateAccount :one
insert into account (username, password) values ($1, $2) returning *;

-- name: GetAccount :one
SELECT * FROM account
WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM account;


-- name: ListOrderedAccounts :many
SELECT * FROM account ORDER BY username OFFSET $1 LIMIT $2;