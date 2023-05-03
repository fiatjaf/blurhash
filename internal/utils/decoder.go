package utils

import (
	"fmt"
	"github.com/bbrks/go-blurhash"
	"image/png"
	"os"
)

func Decode(
	hash string,
	width int,
	height int,
	punch int,
	output string,
) (err error) {
	result, err := blurhash.Decode(hash, width, height, punch)
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.Create(output)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Encode to `PNG` with `DefaultCompression` level
	// then save to file
	err = png.Encode(file, result)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
