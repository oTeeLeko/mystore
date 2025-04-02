-- name: AddCustomer :exec
INSERT INTO Customers (
  FirstName,
  LastName,
  Gender,
  Tel,
  Email_Address
) VALUES (
  ?, ?, ?, ?, ?
);

-- name: GetCustomer :one
SELECT * FROM Customers
WHERE ID = ?
LIMIT 1;

-- name: GetListCustomers :many
SELECT * FROM Customers
LIMIT ? OFFSET ?;

-- name: DeleteCustomer :exec
DELETE FROM Customers
WHERE ID = ?;

-- name: UpdateCustomer :exec
UPDATE Customers
SET FirstName = ?,
    LastName = ?,
    Gender = ?,
    Tel = ?,
    Email_Address = ?,
    Modified = current_timestamp
WHERE ID = ?;