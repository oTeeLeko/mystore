package tbl

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	Name  *string          `gorm:"size:255;notnull"`
	Price *decimal.Decimal `gorm:"type:decimal(10,2)"`
}
