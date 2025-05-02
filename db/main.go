package main

import (
	"financialmanager/db/db"
	"financialmanager/db/web"
	"fmt"
	"log"
)

func main() {
	log.Print("Starting application...")

	//	Инициализация базы данных
	if err := db.Init(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer func() {
		db.CloseDB()
		log.Print("db closed")
	}()

	// Инициализация эндпоинтов
	if err := web.Init(); err != nil {
		fmt.Printf("Web init failed: %v\n", err)
		return
	}

	//	Прослушка
	if err := web.Listening(); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
