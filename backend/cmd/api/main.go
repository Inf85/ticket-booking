package main

import (
	"fmt"
	"github.com/Inf85/ticket-booking/config"
	"github.com/Inf85/ticket-booking/db"
	"github.com/Inf85/ticket-booking/handlers"
	"github.com/Inf85/ticket-booking/repository"
	"github.com/Inf85/ticket-booking/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      "TicketBooking",
		ServerHeader: "Fiber",
	})

	//Repositories
	eventRepository := repository.NewEventRepository(db)
	ticketRepository := repository.NewTicketRepository(db)
	authRepository := repository.NewAuthRepository(db)

	//Services
	authService := services.NewAuthService(authRepository)
	//Routing
	server := app.Group("/api")

	handlers.NewAuthHandler(server.Group("/auth"), authService)
	privateRoutes := server.Use(middleware.AuthProtected(db))
	//Handlers

	handlers.NewEventHandler(server.Group("/event"), eventRepository)
	handlers.NewTicketHandler(server.Group("/ticket"), ticketRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}
