package router

import (
	"github.com/gin-gonic/gin"
	"github.com/orgmange/order-service/internal/handler"
)

func SetupRouter(healthHandler handler.HealthHandler,
	userHandler handler.UserHandler,
	orderHandler handler.OrderHandler,
) *gin.Engine {
	r := gin.Default()
	r.GET("/health", healthHandler.HandleHealth)

	r.GET("/users/:id", userHandler.GetUser)
	r.PUT("/users/:id", userHandler.UpdateUser)
	r.POST("/users", userHandler.CreateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)

	r.GET("/orders/:id", orderHandler.GetOrder)
	r.PUT("/orders/:id", orderHandler.UpdateOrder)
	r.POST("/orders", orderHandler.CreateOrder)
	r.DELETE("/orders/:id", orderHandler.DeleteOrder)
	return r
}
