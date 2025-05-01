package main

import (
	"financialmanager/db/db"
	"financialmanager/db/web"
	"fmt"
)

func main() {
	fmt.Println("Starting application...")

	//	Инициализация базы данных
	db.Initdb()

	// Инициализация эндпоинтов
	if err := web.Init(); err != nil {
		fmt.Printf("Web init failed: %v\n", err)
		return
	}

	//	Прослушка
	fmt.Println("Starting server (blocking mode)...")
	if err := web.Listening(); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
