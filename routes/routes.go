package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/khanjaved9700/orders/order"
	"github.com/khanjaved9700/orders/payment"
)

func RegisterRoutes(r *gin.Engine, orderHandler *order.Handler, paymentHandler *payment.Handler) {
	v1 := r.Group("/api/v1")
	{
		v1.POST("/orders", orderHandler.Create)
		v1.GET("/orders/:id", orderHandler.Get)

		v1.POST("/payments", paymentHandler.Create)
		v1.GET("/payments/:id", paymentHandler.Get)
	}
}
