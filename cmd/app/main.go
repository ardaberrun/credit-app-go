package main

import (
	"github/ardaberrun/credit-app-go/internal/app/handler"
	middleware "github/ardaberrun/credit-app-go/internal/app/middleware/auth"
	"github/ardaberrun/credit-app-go/internal/app/repository"
	"github/ardaberrun/credit-app-go/internal/app/service"
	"github/ardaberrun/credit-app-go/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.Connect();
	if err != nil {
		panic(err);
	}

	router := gin.Default();
	router.Use(middleware.AuthMiddleware());

	userRepository := repository.InitializeUserRepository(db);
	userService := service.InitializeUserService(userRepository);
	userHandler := handler.InitializeUserHandler(userService, router)

	transactionRepository := repository.InitializeTransactionRepository(db);
	transactionService := service.InitializeTransactionService(transactionRepository)
	transactionHandler := handler.InitializeTransactionHandler(transactionService, router)

	userHandler.RegisterRoutes();
	transactionHandler.RegisterRoutes();

	router.Run(":8080");
}