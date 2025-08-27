package order

import "time"

type CreateOrderRequest struct {
	Amount float64 `json:"amount" validate:"required"`
	Status string  `json:"status,omitempty"`
}

type OrderResponse struct {
	ID        uint      `json:"id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

