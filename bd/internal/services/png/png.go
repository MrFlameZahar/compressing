package services

import (
	"fmt"
	"image"
	"image/jpeg"
	"net/http"
	"os"

	"github.com/nfnt/resize"
)

type PngServices struct{}

func (p *PngServices) Decode(inputFile string) (image.Image, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		panic(err)
	}

	mimeType := http.DetectContentType(buffer)

	if mimeType != "image/png" {
		return nil, fmt.Errorf("формат не пнг")
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("something wrong")
	}
	return img, nil
}

func (p *PngServices) Resize(img image.Image) image.Image {
	return resize.Resize(0, 500, img, resize.Lanczos3)
}

func (p *PngServices) Compress(img image.Image, resultName string, quality int) error {

	outFile, err := os.Create(resultName + ".jpg")
	if err != nil {
		return err
	}
	defer outFile.Close()

	jpeg.Encode(outFile, img, &jpeg.Options{Quality: quality})

	return nil
}
