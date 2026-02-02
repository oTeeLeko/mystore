package tbl

import "gorm.io/gorm"

type Inventory struct {
	gorm.Model
	ProductID *uint
	Quantity  *int
}
