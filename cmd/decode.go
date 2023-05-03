package cmd

import (
	"fmt"
	"github.com/bbrks/go-blurhash"
	"github.com/spf13/cobra"
	"image/png"
	"os"
)

var (
	hash   string
	width  int
	height int
	punch  int
	output string
)

func init() {
	decodeCmd.Flags().StringVar(&hash, "hash", "", "String is represented BlurHash code")
	decodeCmd.Flags().StringVarP(&output, "output", "o", "", "Path for the output image")
	decodeCmd.Flags().IntVar(&width, "width", 64, "Width of the output image")
	decodeCmd.Flags().IntVar(&height, "height", 64, "Height of the output image")
	decodeCmd.Flags().IntVar(&punch, "punch", 1, "Intensity of colors in resulted image")

	decodeCmd.MarkFlagRequired("hash")
	decodeCmd.MarkFlagRequired("output")

	decodeCmd.MarkFlagsRequiredTogether("hash", "output")
	decodeCmd.MarkFlagsRequiredTogether("width", "height")

	rootCmd.AddCommand(decodeCmd)
}

var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decode image",
	Long:  `Decode provided blurhash string into image`,
	Run: func(cmd *cobra.Command, args []string) {
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
	},
}
