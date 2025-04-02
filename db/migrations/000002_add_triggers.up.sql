CREATE TRIGGER before_insert_orders
BEFORE INSERT ON Orders
FOR EACH ROW
BEGIN
    DECLARE product_price FLOAT;
    SELECT Price INTO product_price FROM Products WHERE ID = NEW.ProductID;
    SET NEW.Amount = NEW.Quantity * product_price;
END;