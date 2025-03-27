CREATE TABLE `Customers` (
  `ID` varchar(36) PRIMARY KEY DEFAULT UUID(),
  `FirstName` varchar(255) NOT NULL,
  `LastName` varchar(255) NOT NULL, 
  `Gender` varchar(10) NOT NULL,
  `Tel` varchar(10) NOT NULL,
  `Email_Address` varchar(255) UNIQUE NOT NULL,
  `Created` timestamp NOT NULL DEFAULT current_timestamp,
  `Modified` timestamp NOT NULL DEFAULT current_timestamp
);

CREATE TABLE `Products` (
    `ID` varchar(36) PRIMARY KEY DEFAULT UUID(),
    `Name` varchar(255) NOT NULL,
    `Price` float NOT NULL,
    `Created` timestamp NOT NULL DEFAULT current_timestamp,
    `Modified` timestamp NOT NULL DEFAULT current_timestamp
);

CREATE TABLE `Inventories` (
    `ID` varchar(36) PRIMARY KEY DEFAULT UUID(),
    `ProductID` varchar(36) NOT NULL,
    `Quantity` int NOT NULL,
    `Created` timestamp NOT NULL DEFAULT current_timestamp,
    `Modified` timestamp NOT NULL DEFAULT current_timestamp
);

CREATE TABLE `Orders` (
    `ID` varchar(36) PRIMARY KEY DEFAULT UUID(),
    `CustomerID` varchar(36) NOT NULL,
    `ProductID` varchar(36) NOT NULL,
    `Quantity` int NOT NULL,
    `Amount` float NOT NULL,
    `Created` timestamp NOT NULL DEFAULT current_timestamp,
    `Modified` timestamp NOT NULL DEFAULT current_timestamp
);