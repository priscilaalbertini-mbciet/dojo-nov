package entities

import "time"

type User struct {
	ID        int64
	Orders    []Order `gorm:"foreignKey:UserId"`
	Name      string
	CreatedAt time.Time
	UpdateAt  time.Time
}
