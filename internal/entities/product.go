package entities

import (
	"time"

	"github.com/shopspring/decimal"
)

type Product struct {
	ID          int64
	Name        string
	Description string
	Price       decimal.Decimal
	Quantity    decimal.Decimal
	Total       decimal.Decimal
	Orders      []*Order `gorm:"many2many:order_products;"`
	CreatedAt   time.Time
	UpdateAt    time.Time
}

// ID já é considerado primary key
