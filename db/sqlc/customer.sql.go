// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: customer.sql

package db

import (
	"context"
)

const addCustomer = `-- name: AddCustomer :exec
INSERT INTO Customers (
  FirstName,
  LastName,
  Gender,
  Tel,
  Email_Address
) VALUES (
  ?, ?, ?, ?, ?
)
`

type AddCustomerParams struct {
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Gender       string `json:"gender"`
	Tel          string `json:"tel"`
	EmailAddress string `json:"email_address"`
}

func (q *Queries) AddCustomer(ctx context.Context, arg AddCustomerParams) error {
	_, err := q.db.ExecContext(ctx, addCustomer,
		arg.Firstname,
		arg.Lastname,
		arg.Gender,
		arg.Tel,
		arg.EmailAddress,
	)
	return err
}

const deleteCustomer = `-- name: DeleteCustomer :exec
DELETE FROM Customers
WHERE ID = ?
`

func (q *Queries) DeleteCustomer(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteCustomer, id)
	return err
}

const getCustomer = `-- name: GetCustomer :one
SELECT id, firstname, lastname, gender, tel, email_address, created, modified FROM Customers
WHERE ID = ?
LIMIT 1
`

func (q *Queries) GetCustomer(ctx context.Context, id string) (Customer, error) {
	row := q.db.QueryRowContext(ctx, getCustomer, id)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.Firstname,
		&i.Lastname,
		&i.Gender,
		&i.Tel,
		&i.EmailAddress,
		&i.Created,
		&i.Modified,
	)
	return i, err
}

const getListCustomers = `-- name: GetListCustomers :many
SELECT id, firstname, lastname, gender, tel, email_address, created, modified FROM Customers
LIMIT ? OFFSET ?
`

type GetListCustomersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetListCustomers(ctx context.Context, arg GetListCustomersParams) ([]Customer, error) {
	rows, err := q.db.QueryContext(ctx, getListCustomers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Customer{}
	for rows.Next() {
		var i Customer
		if err := rows.Scan(
			&i.ID,
			&i.Firstname,
			&i.Lastname,
			&i.Gender,
			&i.Tel,
			&i.EmailAddress,
			&i.Created,
			&i.Modified,
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

const updateCustomer = `-- name: UpdateCustomer :exec
UPDATE Customers
SET FirstName = ?,
    LastName = ?,
    Gender = ?,
    Tel = ?,
    Email_Address = ?,
    Modified = current_timestamp
WHERE ID = ?
`

type UpdateCustomerParams struct {
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Gender       string `json:"gender"`
	Tel          string `json:"tel"`
	EmailAddress string `json:"email_address"`
	ID           string `json:"id"`
}

func (q *Queries) UpdateCustomer(ctx context.Context, arg UpdateCustomerParams) error {
	_, err := q.db.ExecContext(ctx, updateCustomer,
		arg.Firstname,
		arg.Lastname,
		arg.Gender,
		arg.Tel,
		arg.EmailAddress,
		arg.ID,
	)
	return err
}
