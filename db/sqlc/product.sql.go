// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: product.sql

package db

import (
	"context"
)

const addProduct = `-- name: AddProduct :exec
INSERT INTO Products (
    Name,
    Price
) VALUES(
    ?, ?
)
`

type AddProductParams struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (q *Queries) AddProduct(ctx context.Context, arg AddProductParams) error {
	_, err := q.db.ExecContext(ctx, addProduct, arg.Name, arg.Price)
	return err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM Products
WHERE ID = ?
`

func (q *Queries) DeleteProduct(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteProduct, id)
	return err
}

const getListProducts = `-- name: GetListProducts :many
SELECT id, name, price, created, modified FROM Products
LIMIT ? OFFSET ?
`

type GetListProductsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetListProducts(ctx context.Context, arg GetListProductsParams) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, getListProducts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Price,
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

const getProduct = `-- name: GetProduct :one
SELECT id, name, price, created, modified FROM Products
WHERE ID = ?
LIMIT 1
`

func (q *Queries) GetProduct(ctx context.Context, id string) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProduct, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		&i.Created,
		&i.Modified,
	)
	return i, err
}

const updateProduct = `-- name: UpdateProduct :exec
UPDATE Products
SET Name = ?,
    Price = ?,
    Modified = current_timestamp
WHERE ID = ?
`

type UpdateProductParams struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	ID    string  `json:"id"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) error {
	_, err := q.db.ExecContext(ctx, updateProduct, arg.Name, arg.Price, arg.ID)
	return err
}
