package web

import (
	"github.com/gofiber/fiber/v2"
	"financialmanager/db/web/handlers"
)

var app *fiber.App

func Init() error {

	app = fiber.New()

	api := app.Group("/database")

	api.Get("/goals", handlers.TestGoals1)
	api.Post("/goals", handlers.TestGoals2)
	api.Put("/goals", handlers.TestGoals3)

	api.Get("/groups", handlers.TestGroups1)
	api.Post("/groups", handlers.TestGroups2)
	api.Put("/groups", handlers.TestGroups3)
	
	api.Get("/transactions", handlers.TestTransactions1)
	api.Post("/transactions", handlers.TestTransactions2)
	api.Put("/transactions", handlers.TestTransactions2)


	return nil
}

func Listening() error {
	return app.Listen(":3000")
}


