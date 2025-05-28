package handlers

import (
	"fmt"

	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// handlers
// indexHandler возвращает index.html
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Чтение файла index.html
	data, err := os.ReadFile("index.html")
	if err != nil {
		log.Printf("Ошибка при чтении index.html: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	// Парсим форму
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, "Ошибка при парсинге формы", http.StatusBadRequest)
		return
	}

	// Получаем файл
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Ошибка при получении файла", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Чтение содержимого файла
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка при чтении файла", http.StatusInternalServerError)
		return
	}

	// Передача в функцию автоопределения из пакета service
	converted, err := service.AutoDefinition(string(data))
	if err != nil {
		http.Error(w, "Ошибка при создании файла", http.StatusInternalServerError)
		return
	}
	// Генерация имени файла
	timestamp := strings.ReplaceAll(time.Now().UTC().String(), ":", "-")
	ext := filepath.Ext(handler.Filename)
	outputFileName := fmt.Sprintf("converted_%s%s", timestamp, ext)

	// Создание и запись в файл
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		http.Error(w, "Ошибка при создании файла", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	_, err = outputFile.WriteString(converted)
	if err != nil {
		http.Error(w, "Ошибка при записи в файл", http.StatusInternalServerError)
		return
	}

	// Возврат результата клиенту
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(converted))
} //handlers
