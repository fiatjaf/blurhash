package utils

import (
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
		return err
	}

	file, err := os.Create(output)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode to `PNG` with `DefaultCompression` level
	// then save to file
	err = png.Encode(file, result)
	if err != nil {
		return err
	}
	return err
}
