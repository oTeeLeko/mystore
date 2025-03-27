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
