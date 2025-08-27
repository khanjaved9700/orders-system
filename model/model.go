package model

import "time"

type Order struct {
	ID        uint    `gorm:"primaryKey"`
	Amount    float64 `gorm:"not null"`
	Status    string  `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Payment struct {
	ID        uint    `gorm:"primaryKey"`
	OrderID   uint    `gorm:"not null"`
	Amount    float64 `gorm:"not null"`
	Method    string  `gorm:"not null"`
	Status    string  `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
