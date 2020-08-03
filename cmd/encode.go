package cmd

import (
	"fmt"
	"github.com/gzcharleszhang/stego/pkg/stegolsb"
	"github.com/gzcharleszhang/stego/utils"
	"github.com/spf13/cobra"
	"image"
	"image/png"
	"os"
	"strings"
)

const defaultOutputPrefix = "out"

// encodeCmd represents the encode command
var (
	encodeOutputPath string
	encodeData       string
	encodeCmd        = &cobra.Command{
		Use:   "encode",
		Short: "encodes data in an image",
		Long: `
Stego encode will embed your data in an image.
It breaks the data into bits and writes them to
the least significant bit of each pixel's RGB channel
in the image.`,
		Args: cobra.ExactArgs(1),
		RunE: encode,
	}
)

func writeImage(img image.Image, format, originalPath, defaultOutputPrefix string) error {
	if encodeOutputPath == "" {
		originalPathName, ext := utils.PathWithExtension(originalPath)
		outPath := strings.Join([]string{originalPathName, defaultOutputPrefix}, "-")
		encodeOutputPath = strings.Join([]string{outPath, ext}, ".")
	}
	writer, err := os.Create(encodeOutputPath)
	if writer != nil {
		defer writer.Close()
		defer writer.Sync()
	}
	if err != nil {
		return fmt.Errorf("error creating output image: %v", err)
	}
	if format == "png" {
		err = png.Encode(writer, img)
	}
	if err != nil {
		return fmt.Errorf("error encoding output image: %v", err)
	}
	return nil
}

func encode(cmd *cobra.Command, args []string) error {
	imagePath := args[0]
	img, format, err := utils.GetImage(imagePath)
	if err != nil {
		return err
	}
	outImg, err := stegolsb.LSBEncode(img, encodeData)
	if err != nil {
		return fmt.Errorf("error encoding message: %v", err)
	}
	err = writeImage(outImg, format, imagePath, defaultOutputPrefix)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(encodeCmd)

	encodeCmd.Flags().StringVarP(&encodeData, "data", "d", "", "Data to be encoded")
	encodeCmd.MarkFlagRequired("message")
	encodeCmd.Flags().StringVarP(&encodeOutputPath, "output", "o", "", "Output path for the encoded image")
}
