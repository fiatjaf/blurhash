package cmd

import (
	"fmt"
	"github.com/bbrks/go-blurhash"
	"github.com/spf13/cobra"
	"image/png"
	"os"
)

var (
	blurHash string
	width    int
	height   int
	punch    int
	output   string
)

func init() {
	decodeCmd.Flags().StringVarP(&blurHash, "blurhash", "b", "", "")
	decodeCmd.Flags().IntVar(&width, "width", 0, "")
	decodeCmd.Flags().IntVar(&height, "height", 0, "")
	decodeCmd.Flags().IntVarP(&punch, "punch", "p", 0, "")
	decodeCmd.Flags().StringVarP(&output, "output", "o", "", "")

	decodeCmd.MarkFlagRequired("blurhash")
	decodeCmd.MarkFlagRequired("width")
	decodeCmd.MarkFlagRequired("height")
	decodeCmd.MarkFlagRequired("punch")
	decodeCmd.MarkFlagRequired("output")

	decodeCmd.MarkFlagsRequiredTogether("blurhash", "width", "height", "punch", "output")

	rootCmd.AddCommand(decodeCmd)
}

var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decode image",
	Long:  `Decode provided blurhash string into image`,
	Run: func(cmd *cobra.Command, args []string) {
		result, err := blurhash.Decode(blurHash, width, height, punch)
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
