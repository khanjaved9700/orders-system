package payment

import "time"

// Request payload when client creates a payment
type CreatePaymentRequest struct {
	OrderID uint    `json:"order_id" binding:"required"`
	Amount  float64 `json:"amount" binding:"required"`
	Method  string  `json:"method" binding:"required"` // e.g., CARD, UPI, NETBANKING
}

// Response payload returned to client
type PaymentResponse struct {
	ID        uint      `json:"id"`
	OrderID   uint      `json:"order_id"`
	Amount    float64   `json:"amount"`
	Method    string    `json:"method"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
