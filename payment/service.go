package payment

import (
	"encoding/json"
	"errors"

	"github.com/khanjaved9700/orders/kafka"
	"github.com/khanjaved9700/orders/order"
	"gorm.io/gorm"
)

type Service interface {
	ProcessPayment(payment *CreatePaymentRequest) (*PaymentResponse, error)
	GetPayment(id uint) (*PaymentResponse, error)
}

type service struct {
	repo      Repository
	orderRepo order.Repository
	producer  kafka.Producer
	db        *gorm.DB
}

func NewService(r Repository, or order.Repository, prod kafka.Producer, db *gorm.DB) Service {
	return &service{repo: r, orderRepo: or, producer: prod, db: db}
}

func (s *service) ProcessPayment(req *CreatePaymentRequest) (*PaymentResponse, error) {
	var resp *PaymentResponse

	// run in DB transaction
	err := s.db.Transaction(func(tx *gorm.DB) error {
		// check order exists
		o, err := s.orderRepo.GetByID(req.OrderID)
		if err != nil {
			return errors.New("order not found")
		}

		// create payment
		payment, err := s.repo.Create(req)
		if err != nil {
			return err
		}

		// update order status to PAID
		o.Status = "PAID"
		if err := tx.Save(o).Error; err != nil {
			return err
		}

		// publish Kafka event with persisted payment
		payload, _ := json.Marshal(payment)
		_ = s.producer.Publish(kafka.TopicOrderEvents, string(payload))

		resp = &PaymentResponse{
			ID:        payment.ID,
			OrderID:   payment.OrderID,
			Amount:    payment.Amount,
			Method:    payment.Method,
			Status:    payment.Status,
			CreatedAt: payment.CreatedAt,
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) GetPayment(id uint) (*PaymentResponse, error) {
	return s.repo.GetByID(id)
}
