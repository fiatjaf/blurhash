package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Flags for encode command
var (
	componentX int
	componentY int
	input      string
	autodetect bool
)

// FLags for decode command
var (
	hash   string
	width  int
	height int
	punch  int
	output string
)

// Commands descriptions
var (
	rootCmd = &cobra.Command{
		Short: "BlurHash CLI tool",
		Long:  "CLI tool for the working with BlurHash",
	}

	encodeCmd = &cobra.Command{
		Use:   "encode",
		Short: "Encode image",
		Long:  `Encode provided image in the blurhash string`,
		Run: func(cmd *cobra.Command, args []string) {
			var hash string
			var err error
			if autodetect == false {
				hash, err = Encode(componentX, componentY, input)
			} else {
				hash, err = EncodeA(input)
			}
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			} else {
				fmt.Fprintln(os.Stdout, hash)
				os.Exit(0)
			}
		},
	}

	decodeCmd = &cobra.Command{
		Use:   "decode",
		Short: "Decode image",
		Long:  `Decode provided blurhash string into image`,
		Run: func(cmd *cobra.Command, args []string) {
			err := Decode(hash, width, height, punch, output)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			} else {
				fmt.Fprintln(os.Stdout, "OK")
				os.Exit(0)
			}
		},
	}
)

func init() {
	cobra.OnInitialize()
	initEncode()
	initDecode()
}

func initEncode() {
	encodeCmd.Flags().IntVarP(&componentX, "componentX", "x", 9, "Quantity of components by X axis")
	encodeCmd.Flags().IntVarP(&componentY, "componentY", "y", 9, "Quantity of components by Y axis")
	encodeCmd.Flags().StringVarP(&input, "input", "i", "", "Path for the input image")
	encodeCmd.Flags().BoolVarP(&autodetect, "autodetect", "a", false, "Autodetect X/Y components numbers")

	encodeCmd.MarkFlagRequired("input")

	encodeCmd.MarkFlagsRequiredTogether("componentX", "componentY")

	rootCmd.AddCommand(encodeCmd)
}

func initDecode() {
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
