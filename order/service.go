package order

import (
	"encoding/json"
	"fmt"

	"github.com/khanjaved9700/orders/kafka"
	"github.com/khanjaved9700/orders/model"
	"github.com/khanjaved9700/orders/redis"
)

type Service interface {
	CreateOrder(order *CreateOrderRequest) (OrderResponse, error)
	GetOrder(id uint) (OrderResponse, error)
}

type service struct {
	repo  Repository
	prod  kafka.Producer
	cache redis.Cache
}

func NewService(r Repository, p kafka.Producer, c redis.Cache) Service {
	return &service{repo: r, prod: p, cache: c}
}

func (s *service) CreateOrder(req *CreateOrderRequest) (OrderResponse, error) {
	// map request → model
	order := model.Order{
		Amount: req.Amount,
		Status: "PENDING", // default
	}

	// save to DB
	if err := s.repo.Create(&order); err != nil {
		return OrderResponse{}, err
	}

	// map to response
	resp := OrderResponse{
		ID:        order.ID,
		Amount:    order.Amount,
		Status:    order.Status,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	}

	// publish kafka event
	payload, _ := json.Marshal(resp)
	_ = s.prod.Publish(kafka.TopicOrderEvents, string(payload))

	// cache order
	_ = s.cache.Set(fmt.Sprintf("order_%d", order.ID), string(payload))

	return resp, nil
}

func (s *service) GetOrder(id uint) (OrderResponse, error) {
	// try redis first
	val, err := s.cache.Get(fmt.Sprintf("order_%d", id))
	if err == nil && val != "" {
		var o OrderResponse
		if err := json.Unmarshal([]byte(val), &o); err == nil {
			return o, nil
		}
	}

	// fallback to DB
	o, err := s.repo.GetByID(id)
	if err != nil {
		return OrderResponse{}, err
	}

	// map model → response
	return OrderResponse{
		ID:        o.ID,
		Amount:    o.Amount,
		Status:    o.Status,
		CreatedAt: o.CreatedAt,
		UpdatedAt: o.UpdatedAt,
	}, nil
}
