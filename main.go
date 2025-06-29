package main

import (
	"Smarket/api_handlers"
	_ "Smarket/docs"
	"Smarket/internal/configs"
	"Smarket/internal/db"
	"Smarket/pkg/logger"
	"log"
)

// @title Smarket API
// @version 1.0
// @description This is the Smarket API for managing sales and inventory
// @host localhost:8050
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	if err := configs.ReadSettings(); err != nil {
		log.Fatalf("Ошибка чтения настроек: %s", err)
	}

	if err := logger.Init(); err != nil {
		log.Fatalf("Ошибка инициализации логгера: %s", err)
	}
	logger.Info.Println("Loggers initialized successfully!")

	if err := db.ConnectDB(); err != nil {
		logger.Error.Printf("Ошибка подключения к БД: %s", err.Error())
		return
	}
	logger.Info.Println("Connection to database established successfully!")

	if err := api_handlers.RunServer(); err != nil {
		logger.Error.Printf("Error during running HTTP server: %s", err.Error())
		return
	}
}
