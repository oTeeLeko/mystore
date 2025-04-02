-- name: AddProduct :exec
INSERT INTO Products (
    Name,
    Price
) VALUES(
    ?, ?
);

-- name: GetProduct :one
SELECT * FROM Products
WHERE ID = ?
LIMIT 1;

-- name: GetListProducts :many
SELECT * FROM Products
LIMIT ? OFFSET ?;

-- name: DeleteProduct :exec
DELETE FROM Products
WHERE ID = ?;

-- name: UpdateProduct :exec
UPDATE Products
SET Name = ?,
    Price = ?,
    Modified = current_timestamp
WHERE ID = ?;