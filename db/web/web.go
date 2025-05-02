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
	api.Post("/goal", handlers.TestGoals2)
	api.Put("/goal", handlers.TestGoals3)
	api.Delete("/goal", handlers.TestGoals3)

	api.Get("/groups", handlers.TestGroups1)
	api.Post("/group", handlers.TestGroups2)
	api.Put("/group", handlers.TestGroups3)
	api.Delete("/group", handlers.TestGroups3)
	
	api.Get("/transactions", handlers.GetTransactions)			//	Получение списка транзакций
	api.Post("/transaction", handlers.AddTransactions)			//	Добавление новой транзакции
	api.Put("/transaction", handlers.UpdateTransactions)		//	Обновление информации о транзакции
	api.Delete("/transaction", handlers.DeleteTransactions)		//	Удаление транзакции

	return nil
}

func Listening() error {
	log.Println("Listening on 3000 port...")
	return app.Listen(":3000")
}


