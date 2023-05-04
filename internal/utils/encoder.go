package utils

import (
	"github.com/bbrks/go-blurhash"
	"image"
	"math"
	"os"
)

func Encode(
	componentX int,
	componentY int,
	input string,
) (hash string, err error) {
	file, err := os.Open(input)
	if err != nil {
		return hash, err
	}
	img, _, err := image.Decode(file)
	if err != nil {
		return hash, err
	}

	hash, err = blurhash.Encode(componentX, componentY, img)
	if err != nil {
		return hash, err
	}
	return hash, err
}

func EncodeA(input string) (hash string, err error) {
	file, err := os.Open(input)
	if err != nil {
		return hash, err
	}
	img, _, err := image.Decode(file)
	if err != nil {
		return hash, err
	}

	var componentX, componentY int
	width, height := img.Bounds().Dx(), img.Bounds().Dy()
	if ratio := float64(width) / float64(height); width >= height {
		componentX = 9
		componentY = int(math.Min(math.Round(float64(componentX)/ratio), 9))
	} else {
		componentY = 9
		componentX = int(math.Min(math.Round(float64(componentY)*ratio), 9))
	}

	hash, err = blurhash.Encode(componentX, componentY, img)
	if err != nil {
		return hash, err
	}
	return hash, err
}
