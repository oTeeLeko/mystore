-- name: AddOrder :exec
INSERT INTO Orders (
  CustomerID,
  ProductID,
  Quantity
) VALUES (
  ?, ?, ?
);

-- name: GetOrder :one
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
LIMIT 1;

-- name: GetListOrders :many
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
LIMIT ? OFFSET ?;

-- name: DeleteOrder :exec
DELETE FROM Orders
WHERE ID = ?;