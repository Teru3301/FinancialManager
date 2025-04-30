package web

import (
	"financialmanager/db/web/handlers"
	"fmt"
	"net/http"
)

// регистрирует обработчики, но не запускает сервер
func Init() error {
	fmt.Println("Initializing web routes...")

	http.HandleFunc("/transactions", func(w http.ResponseWriter, r *http.Request) {
		handlers.Transtestfunc(w, r)
	})

	http.HandleFunc("/goals", func(w http.ResponseWriter, r *http.Request) {
		handlers.Goalstestfunc(w, r)
	})

	http.HandleFunc("/groups", func(w http.ResponseWriter, r *http.Request) {
		handlers.Groupstestfunc(w, r)
	})

	return nil
}

// Начало прослушки эндпоинтов
func Listening() error {
	port := ":8080"
	fmt.Println("Server is starting on " + port + "...")
	return http.ListenAndServe(port, nil)
}
