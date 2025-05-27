package main

import (
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
	//"internal/server"
)

func main() {
	logger := log.New(log.Writer(), "HTTP Server: ", log.LstdFlags)
	server := server.NewServer(logger)

	logger.Println("Запуск сервера")
	err := server.HTTPServer.ListenAndServe()
	if err != nil {
		logger.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
