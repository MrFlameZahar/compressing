package services

import "image"

type CompressingService interface {
	Decode(inputFile string) (image.Image, error)
	Resize(img image.Image) image.Image
	Compress(img image.Image, resultName string, quality int) error
}
