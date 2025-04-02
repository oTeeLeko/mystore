-- name: AddInventory :exec
INSERT INTO Inventories (
  ProductID,
  Quantity
) VALUES (
  ?, ?
);

-- name: GetInventory :one
SELECT 
  Inventories.ID,
  Inventories.ProductID,
  Products.Name,
  Inventories.Quantity
FROM Inventories
LEFT JOIN Products ON Inventories.ProductID = Products.ID
WHERE Inventories.ID = ?
LIMIT 1;

-- name: GetListInventories :many
SELECT 
  Inventories.ID,
  Inventories.ProductID,
  Products.Name,
  Inventories.Quantity
FROM Inventories
LEFT JOIN Products ON Inventories.ProductID = Products.ID
LIMIT ? OFFSET ?;

-- name: DeleteInventory :exec
DELETE FROM Inventories
WHERE ID = ?;

-- name: UpdateInventoryQuantity :exec
UPDATE Inventories
SET Quantity = Quantity + ?,
    Modified = current_timestamp
WHERE ProductID = ?;