package web

import (
	"github.com/gofiber/fiber/v2"
	"financialmanager/db/web/handlers"
	"log"
)

var app *fiber.App

func Init() error {
	log.Println("Initializing routes...")
	
	app = fiber.New()

	api := app.Group("/database")

	api.Get("/goals", handlers.TestGoals1)
	api.Post("/goals", handlers.TestGoals2)
	api.Put("/goals", handlers.TestGoals3)

	api.Get("/groups", handlers.TestGroups1)
	api.Post("/groups", handlers.TestGroups2)
	api.Put("/groups", handlers.TestGroups3)
	
	api.Get("/transactions", handlers.GetTransactions)			//	Получение списка транзакций
	api.Post("/transactions", handlers.AddTransactions)			//	Добавление новой транзакции
	api.Put("/transactions", handlers.UpdateTransactions)		//	Обновление информации о транзакции
	api.Delete("/transactions", handlers.DeleteTransactions)	//	Удаление транзакции

	return nil
}

func Listening() error {
	log.Println("Listening on 3000 port...")
	return app.Listen(":3000")
}


