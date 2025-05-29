package server

import (
	//_ "internal/handlers"
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	_ "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

type Server struct {
	Logger     *log.Logger
	HTTPServer *http.Server
}

// NewServer новый http-сервер
func NewServer(logger *log.Logger) *Server {
	// http-роутер
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

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
