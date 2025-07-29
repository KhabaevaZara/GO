package main

import (
	"log"
	"my-go-backend/database"
	"my-go-backend/handlers"
	"net/http"
)

func main() {
	// Инициализация БД
	if err := database.InitDB(); err != nil {
		log.Fatal("Database init failed:", err)
	}

	// Маршруты API
	http.HandleFunc("/api/tasks", handlers.GetTasks)
	http.HandleFunc("/api/tasks/create", handlers.CreateTask)
	http.HandleFunc("/api/tasks/delete", handlers.DeleteTask)

	// Статические файлы
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// Запуск сервера
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
