package tbl

import "gorm.io/gorm"

type Orders struct {
	gorm.Model
	CustomerID uint
	ProductID  uint
	Quantity   int32
}
