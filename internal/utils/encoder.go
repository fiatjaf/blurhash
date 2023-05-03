package utils

import (
	"fmt"
	"github.com/bbrks/go-blurhash"
	"image"
	"math"
	"os"
)

func Encode(
	componentX int,
	componentY int,
	input string,
) (err error) {
	file, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := blurhash.Encode(componentX, componentY, img)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
	return
}

func EncodeA(input string) (err error) {
	file, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println(err)
		return
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

	result, err := blurhash.Encode(componentX, componentY, img)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
	return
}
