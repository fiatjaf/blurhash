package cmd

import (
	"fmt"
	"github.com/bbrks/go-blurhash"
	"github.com/spf13/cobra"
	"image"
	_ "image/png"
	"os"
)

var (
	componentX int
	componentY int
	input      string
)

func init() {
	encodeCmd.Flags().IntVarP(&componentX, "componentX", "x", 0, "")
	encodeCmd.Flags().IntVarP(&componentY, "componentY", "y", 0, "")
	encodeCmd.Flags().StringVarP(&input, "input", "i", "", "")

	encodeCmd.MarkFlagRequired("componentX")
	encodeCmd.MarkFlagRequired("componentY")
	encodeCmd.MarkFlagRequired("input")

	encodeCmd.MarkFlagsRequiredTogether("componentX", "componentY", "input")

	rootCmd.AddCommand(encodeCmd)
}

var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Encode image",
	Long:  `Encode provided image in the blurhash string`,
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.Open(input)
		if err != nil {
			fmt.Println(err)
			return
		}
		image, _, err := image.Decode(file)
		if err != nil {
			fmt.Println(err)
			return
		}

		result, err := blurhash.Encode(componentX, componentY, image)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(result)
	},
}
