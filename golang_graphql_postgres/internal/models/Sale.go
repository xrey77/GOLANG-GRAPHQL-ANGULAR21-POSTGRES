package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Sale struct {
	gorm.Model
	ID         int             `gorm:"type:integer;primarykey"`
	Saleamount decimal.Decimal `gorm:"type:numeric(10,2);default:0.00"`
	Salesdate  time.Time       `gorm:"default:CURRENT_TIMESTAMP(3)"`
}
