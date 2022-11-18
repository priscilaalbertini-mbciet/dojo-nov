package entities

import (
	"time"

	"github.com/shopspring/decimal"
)

type Order struct {
	Id        int64      `gorm:"primaryKey"`
	Items     []*Product `gorm:"many2many:order_products;"`
	User      User       `gorm:"foreignKey:user_id"`
	UserId    int64
	Status    string
	Total     decimal.Decimal
	CreatedAt time.Time
	UpdateAt  time.Time
}

// CreatedAt e UpdatedAt são populados automaticamente com a data atual caso não sejam preenchidos
