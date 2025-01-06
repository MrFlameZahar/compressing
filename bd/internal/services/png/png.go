package services

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/nfnt/resize"
)

type pngServices struct{}

func NewPngServices() *pngServices {
	return &pngServices{}
}

func (p *pngServices) Decode(fileName string) error {

	file, _ := os.Open(fileName)
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("something wrong")
	}

	resizing(&img)
	compress(img)

	return nil
}

func compress(img image.Image) error {

	outFile, err := os.Create("uploaded_image.png")
	if err != nil {
		return err
	}
	defer outFile.Close()

	png.Encode(outFile, img)

	return nil
}

func resizing(img *image.Image) {
	resize.Resize(0, 500, *img, resize.Lanczos3)
}
