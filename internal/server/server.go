package server

import (
	"log"
	"net/http"
	"time"
)


type Server struct {
	Logger *log.Logger
	HTTPServer *http.Server
}

// NewServer новый http-сервер
func NewServer(logger *log.Logger) *Server {
	// http-роутер
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Println("Запрос на корневой маршрут")
		w.Write([]byte("запущен"))
	})

	// Настройка http-сервера
	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger:     logger,
		HTTPServer: httpServer,
	}
}