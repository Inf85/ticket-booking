package main

import (
	"fmt"
	"github.com/Inf85/ticket-booking/config"
	"github.com/Inf85/ticket-booking/db"
	"github.com/Inf85/ticket-booking/handlers"
	"github.com/Inf85/ticket-booking/repository"
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

	//Routing
	server := app.Group("/api")

	//Handlers
	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}