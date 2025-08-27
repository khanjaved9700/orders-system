package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/khanjaved9700/orders/config"
	"github.com/khanjaved9700/orders/kafka"
	"github.com/khanjaved9700/orders/model"
	"github.com/khanjaved9700/orders/order"
	"github.com/khanjaved9700/orders/payment"
	"github.com/khanjaved9700/orders/redis"
	"github.com/khanjaved9700/orders/routes"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Config load failed:", err)
	}
	config.OverrideFromEnv(cfg)

	db, err := config.ConnectDB(cfg)
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	// migrate
	if err := db.AutoMigrate(&model.Order{}, &model.Payment{}); err != nil {
		log.Fatal("AutoMigrate failed:", err)
	}

	// init deps
	cache := redis.NewCache(fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port))
	prod := kafka.NewProducer(strings.Join(cfg.Kafka.Brokers, ","))

	orderRepo := order.NewRepository(db)
	orderSvc := order.NewService(orderRepo, prod, cache)
	orderHandler := order.NewHandler(orderSvc)

	paymentRepo := payment.NewRepository(db)
	paymentSvc := payment.NewService(paymentRepo, orderRepo, prod, db)
	paymentHandler := payment.NewHandler(paymentSvc)

	// gin
	r := gin.Default()
	routes.RegisterRoutes(r, orderHandler, paymentHandler)

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Println("Server running at", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
