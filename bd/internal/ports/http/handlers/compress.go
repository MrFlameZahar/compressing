package handlers

import (
	repo "bd/bd/internal/repo/redis"

	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func Compress(w http.ResponseWriter, r *http.Request) {

	r.Body = http.MaxBytesReader(w, r.Body, 10*1024*1024)

	if r.Header.Get("Content-Type") != "image/png" {
		http.Error(w, "Только файлы PNG поддерживаются", http.StatusUnsupportedMediaType)
		return
	}
	imageID := strconv.Itoa(int(time.Now().Unix()))
	fileName := imageID + ".png"

	outFile, err := os.Create(fileName)

	if err != nil {
		http.Error(w, "Ошибка при создании файла", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, r.Body)
	if err != nil {
		http.Error(w, "Ошибка при записи файла", http.StatusInternalServerError)
		return
	}
	repo.AddToQueue(imageID)

	fmt.Fprintln(w, "Файл успешно загружен")
}

func GetImage(w http.ResponseWriter, r *http.Request) {
	imageID := r.URL.Query().Get("id")
	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(http.StatusOK)

	file, _ := os.Open(imageID + ".png")
	defer file.Close()

	_, _ = io.Copy(w, file)
}
