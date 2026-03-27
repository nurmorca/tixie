package main

import (
	"context"
	"fmt"
	"ticket/common/app"
	"ticket/common/cache"
	"ticket/common/database"
	"ticket/controller"
	"ticket/repository"
	"ticket/service"

	"github.com/labstack/echo/v4"
)

func main() {
	ctx := context.Background()
	e := echo.New()
	configManager := app.NewConfigManager()
	dbPool := database.GetConnectionPool(ctx, configManager.PostgreSqlConfig)
	rdb := cache.GetRedisClient(ctx, configManager.RedisConfig)
	ticketRepository := repository.NewTicketRepository(dbPool)
	ticketService := service.NewTicketService(ticketRepository, rdb)
	productController := controller.NewTicketController(ticketService)
	productController.RegisterRoutes(e)
	for _, r := range e.Routes() {
		fmt.Printf("ROUTE: %s %s\n", r.Method, r.Path)
	}
	e.Start("localhost:8081")
}
