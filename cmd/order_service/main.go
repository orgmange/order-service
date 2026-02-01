package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/orgmange/order-service/internal/config"
	"github.com/orgmange/order-service/internal/handler"
	"github.com/orgmange/order-service/internal/repository"
	"github.com/orgmange/order-service/internal/router"
	"github.com/orgmange/order-service/internal/service"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	healthHandler := handler.NewHealthHandler(config.Version)
	r := router.SetupRouter(*healthHandler, userHandler)
	srv := &http.Server{
		Addr:    config.Address + ":" + config.Port,
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	waitShutdown(srv)
}

func waitShutdown(srv *http.Server) {
	log.Println("server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("server exiting")
}
