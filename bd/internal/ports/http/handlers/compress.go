package handlers

import (
	png "bd/bd/internal/services/png"
	"fmt"
	"io"
	"net/http"
	"os"
)

func Compress(w http.ResponseWriter, r *http.Request) {
	var compressingService = png.NewPngServices()

	r.Body = http.MaxBytesReader(w, r.Body, 10*1024*1024)

	if r.Header.Get("Content-Type") != "image/png" {
		http.Error(w, "Только файлы PNG поддерживаются", http.StatusUnsupportedMediaType)
		return
	}

	outFile, err := os.Create("uploaded_image.png")
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

	compressingService.Decode("uploaded_image.png")

	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(http.StatusOK)

	file, _ := os.Open("uploaded_image.png")
	defer file.Close()

	_, _ = io.Copy(w, file)

	fmt.Fprintln(w, "Файл успешно загружен")
}
