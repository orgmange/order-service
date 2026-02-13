package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/orgmange/order-service/internal/config"
	"github.com/orgmange/order-service/internal/order"
	handler "github.com/orgmange/order-service/internal/transport/http"
	"github.com/orgmange/order-service/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := setupDB(config.DBCfg)
	if err != nil {
		log.Fatal("setuping DB error:", err)
	}

	userRepository := user.NewGormRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	orderRepository := order.NewGormRepository(db)
	orderService := order.NewService(orderRepository)
	orderHandler := handler.NewOrderHandler(orderService)

	healthHandler := handler.NewHealthHandler(config.Version)
	r := handler.SetupRouter(*healthHandler, userHandler, *orderHandler)
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

func setupDB(dbCfg *config.DBCfg) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbCfg.Host,
		dbCfg.User,
		dbCfg.Password,
		dbCfg.Name,
		dbCfg.Port,
		dbCfg.Sslmode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&user.Entity{})
	db.AutoMigrate(&order.Entity{})

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(dbCfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dbCfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(dbCfg.MaxConnsLifetimeSeconds))

	return db, nil
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
