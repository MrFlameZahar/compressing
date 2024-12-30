package main

import (
	"bd/bd/internal/services/png"
	"fmt"
)

func main() {
	pngRealisation := services.NewPngServices()

	img, err := pngRealisation.Decode("newimage.png")

	if err != nil {
		fmt.Printf("Ошибка при декодировании %v", err)
	}

	img = pngRealisation.Resize(img)

	err = pngRealisation.Compress(img, "newimagw2", 80)

	if err != nil {
		fmt.Printf("Ошибка при компрессии%v", err)
	}

	println("Сжатие исполнено")
}
