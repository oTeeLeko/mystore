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