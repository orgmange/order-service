package router

import (
	"github.com/gin-gonic/gin"
	"github.com/orgmange/order-service/internal/config"
	"github.com/orgmange/order-service/internal/handler"
)

func SetupRouter(cfg *config.Config) *gin.Engine {
	r := gin.Default()
	healthHandler := handler.NewHealthHandler(cfg.Version)
	r.GET("/health", healthHandler.HandleHealth)
	return r
}
