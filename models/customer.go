package tbl

import "gorm.io/gorm"

type Customers struct {
	gorm.Model
	Firstname    *string `gorm:"size:255"`
	Lastname     *string `gorm:"size:255"`
	Gender       *string `gorm:"size:10"`
	Tel          *string `gorm:"size:10"`
	EmailAddress *string `gorm:"size:255;uniqueIndex;notnull"`
}
