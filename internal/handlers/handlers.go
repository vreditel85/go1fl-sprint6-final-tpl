package handlers

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	//"strings"
	"time"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

http.HandleFunc("/", indexHandler)
http.HandleFunc("/upload", uploadHandler)

if err := http.ListenAndServe(":8080", nil); err != nil {
fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
}
// indexHandler возвращает index.html
func indexHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Ошибка загрузки", http.StatusInternalServerError)
		return
	}

	if err := templ.Execute(w, nil); err != nil {
		http.Error(w, "Ошибка выполнения", http.StatusInternalServerError)
	}
}

// loadHandler обрабатывает загрузку файла
func loadHandler(w http.ResponseWriter, r *http.Request) {
	//if err := r.ParseMultipartForm(10 << 20); err != nil { // 10 MB
	//	http.Error(w, "Error parsing form", http.StatusInternalServerError)
	//	return
	//}

	file, handler, err := r.FormFile("loadedFile")
	if err != nil {
		http.Error(w, "Ошибка получения файла", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Чтение содержимого файла
	fileData, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
		return
	}

	// Конвертация строки (например, в верхний регистр)
	converted, err := service.AutoDefinition(string(fileData))
	if err != nil {
		http.Error(w, "Ошибка конвертации", http.StatusInternalServerError)
		return
	}

	// Создание нового файла
	ext := filepath.Ext(handler.Filename)
	newFileName := fmt.Sprintf("converted_%s%s", time.Now().UTC().Format("20060102T150405"), ext)

	outFile, err := os.Create(newFileName)
	if err != nil {
		http.Error(w, "Ошибка создания файла", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	// Запись преобразованных данных
	if _, err := outFile.WriteString(converted); err != nil {
		http.Error(w, "Ошибка записи файла", http.StatusInternalServerError)
		return
	}

	// Возврат результата пользователю
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("Converted result:\n\n" + converted))
}