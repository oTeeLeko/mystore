// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: order.sql

package db

import (
	"context"
	"database/sql"
)

const addOrder = `-- name: AddOrder :exec
INSERT INTO Orders (
  CustomerID,
  ProductID,
  Quantity,
  Amount
) VALUES (
  ?, ?, ?, ?
)
`

type AddOrderParams struct {
	Customerid string  `json:"customerid"`
	Productid  string  `json:"productid"`
	Quantity   int32   `json:"quantity"`
	Amount     float64 `json:"amount"`
}

func (q *Queries) AddOrder(ctx context.Context, arg AddOrderParams) error {
	_, err := q.db.ExecContext(ctx, addOrder,
		arg.Customerid,
		arg.Productid,
		arg.Quantity,
		arg.Amount,
	)
	return err
}

const deleteOrder = `-- name: DeleteOrder :exec
DELETE FROM Orders
WHERE ID = ?
`

func (q *Queries) DeleteOrder(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteOrder, id)
	return err
}

const getListOrders = `-- name: GetListOrders :many
SELECT 
    Orders.ID,
    Orders.CustomerID,
    Customers.FirstName,
    Customers.LastName,
    Customers.Tel,
    Customers.Email_Address,
    Orders.ProductID,
    Products.Name,
    Products.Price,
    Orders.Quantity,
    Orders.Amount
FROM Orders
LEFT JOIN Customers ON Orders.CustomerID = Customers.ID
LEFT JOIN Products ON Orders.ProductID = Products.ID
LIMIT ? OFFSET ?
`

type GetListOrdersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetListOrdersRow struct {
	ID           string          `json:"id"`
	Customerid   string          `json:"customerid"`
	Firstname    sql.NullString  `json:"firstname"`
	Lastname     sql.NullString  `json:"lastname"`
	Tel          sql.NullString  `json:"tel"`
	EmailAddress sql.NullString  `json:"email_address"`
	Productid    string          `json:"productid"`
	Name         sql.NullString  `json:"name"`
	Price        sql.NullFloat64 `json:"price"`
	Quantity     int32           `json:"quantity"`
	Amount       float64         `json:"amount"`
}

func (q *Queries) GetListOrders(ctx context.Context, arg GetListOrdersParams) ([]GetListOrdersRow, error) {
	rows, err := q.db.QueryContext(ctx, getListOrders, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetListOrdersRow{}
	for rows.Next() {
		var i GetListOrdersRow
		if err := rows.Scan(
			&i.ID,
			&i.Customerid,
			&i.Firstname,
			&i.Lastname,
			&i.Tel,
			&i.EmailAddress,
			&i.Productid,
			&i.Name,
			&i.Price,
			&i.Quantity,
			&i.Amount,
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

const getOrder = `-- name: GetOrder :one
SELECT 
    Orders.ID,
    Orders.CustomerID,
    Customers.FirstName,
    Customers.LastName,
    Customers.Tel,
    Customers.Email_Address,
    Orders.ProductID,
    Products.Name,
    Products.Price,
    Orders.Quantity,
    Orders.Amount
FROM Orders
LEFT JOIN Customers ON Orders.CustomerID = Customers.ID
LEFT JOIN Products ON Orders.ProductID = Products.ID
WHERE Orders.ID = ?
LIMIT 1
`

type GetOrderRow struct {
	ID           string          `json:"id"`
	Customerid   string          `json:"customerid"`
	Firstname    sql.NullString  `json:"firstname"`
	Lastname     sql.NullString  `json:"lastname"`
	Tel          sql.NullString  `json:"tel"`
	EmailAddress sql.NullString  `json:"email_address"`
	Productid    string          `json:"productid"`
	Name         sql.NullString  `json:"name"`
	Price        sql.NullFloat64 `json:"price"`
	Quantity     int32           `json:"quantity"`
	Amount       float64         `json:"amount"`
}

func (q *Queries) GetOrder(ctx context.Context, id string) (GetOrderRow, error) {
	row := q.db.QueryRowContext(ctx, getOrder, id)
	var i GetOrderRow
	err := row.Scan(
		&i.ID,
		&i.Customerid,
		&i.Firstname,
		&i.Lastname,
		&i.Tel,
		&i.EmailAddress,
		&i.Productid,
		&i.Name,
		&i.Price,
		&i.Quantity,
		&i.Amount,
	)
	return i, err
}
