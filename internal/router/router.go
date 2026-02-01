package router

import (
	"github.com/gin-gonic/gin"
	"github.com/orgmange/order-service/internal/handler"
)

func SetupRouter(healthHandler handler.HealthHandler, userHandler handler.UserHandler) *gin.Engine {
	r := gin.Default()
	r.GET("/health", healthHandler.HandleHealth)

	r.GET("/users/:id", userHandler.GetUser)
	r.PUT("/users/:id", userHandler.UpdateUser)
	r.POST("/users", userHandler.CreateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)
	return r
}
