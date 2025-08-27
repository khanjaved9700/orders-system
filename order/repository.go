package order

import (
	"github.com/khanjaved9700/orders/model"
	"gorm.io/gorm"
)

type Repository interface {
	Create(order *model.Order) error
	GetByID(id uint) (model.Order, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(order *model.Order) error {
	return r.db.Create(order).Error
}

func (r *repository) GetByID(id uint) (model.Order, error) {
	var o model.Order
	if err := r.db.First(&o, id).Error; err != nil {
		return model.Order{}, err
	}
	return o, nil
}
