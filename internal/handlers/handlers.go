package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func ReturnData(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "Файл не найден", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(data)
}
func ReturnConverted(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Не удалось распарсить", http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Не удалось получить файл", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Не удалось прочитать файл", http.StatusInternalServerError)
		return
	}

	result := service.Convert(string(data))

	filename := fmt.Sprintf("%s%s", time.Now().UTC().Format("20060102150405"), filepath.Ext(handler.Filename))

	newFile, err := os.Create(filename)
	if err != nil {
		http.Error(w, "Не удалось создать файл", http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	_, err = newFile.WriteString(result)
	if err != nil {
		http.Error(w, "Не удалось записать файл", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(result))
}
