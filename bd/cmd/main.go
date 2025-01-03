package main

import (
	"bd/bd/internal/services"
	png "bd/bd/internal/services/png"
	"fmt"
)

var compressingService services.CompressingService

func main() {

	compressingService = png.NewPngServices()

	img, err := compressingService.Decode("newimage.png")
	if err != nil {
		fmt.Printf("Ошибка при декодировании: %v\n", err)
		return
	}

	img = compressingService.Resize(img)

	err = compressingService.Compress(img, "newimage2", 80)
	if err != nil {
		fmt.Printf("Ошибка при компрессии: %v\n", err)
		return
	}

	println("Сжатие исполнено")
}
