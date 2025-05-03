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

	api.Get("/goals", handlers.GetGoals)
	api.Post("/goal", handlers.AddGoal)
	api.Put("/goal", handlers.UpdateGoal)
	api.Delete("/goal", handlers.DeleteGoal)

	api.Get("/transactions", handlers.GetTransactions)			//	Получение списка транзакций
	api.Post("/transaction", handlers.AddTransaction)			//	Добавление новой транзакции
	api.Put("/transaction", handlers.UpdateTransaction)			//	Обновление информации о транзакции
	api.Delete("/transaction", handlers.DeleteTransaction)		//	Удаление транзакции

	return nil
}

func Listening() error {
	log.Println("Listening on 3000 port...")
	return app.Listen(":3000")
}


