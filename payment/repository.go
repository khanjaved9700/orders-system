package payment

import (
	"github.com/khanjaved9700/orders/model"
	"gorm.io/gorm"
)

type Repository interface {
	// Create returns the persisted model so caller can access ID/CreatedAt
	Create(req *CreatePaymentRequest) (model.Payment, error)
	Update(payment *model.Payment) error
	GetByID(id uint) (*PaymentResponse, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(req *CreatePaymentRequest) (model.Payment, error) {
	payment := model.Payment{
		OrderID: req.OrderID,
		Amount:  req.Amount,
		Method:  req.Method,
		Status:  "PENDING", // default status on create
	}
	if err := r.db.Create(&payment).Error; err != nil {
		return model.Payment{}, err
	}
	return payment, nil
}

func (r *repository) Update(payment *model.Payment) error {
	return r.db.Save(payment).Error
}

func (r *repository) GetByID(id uint) (*PaymentResponse, error) {
	var payment model.Payment
	if err := r.db.First(&payment, id).Error; err != nil {
		return nil, err
	}
	resp := &PaymentResponse{
		ID:        payment.ID,
		OrderID:   payment.OrderID,
		Amount:    payment.Amount,
		Method:    payment.Method,
		Status:    payment.Status,
		CreatedAt: payment.CreatedAt,
	}
	return resp, nil
}
