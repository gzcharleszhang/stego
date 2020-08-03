package cmd

import (
	"fmt"
	stego_lsb "github.com/gzcharleszhang/stego/pkg/stegolsb"
	"github.com/gzcharleszhang/stego/utils"
	"github.com/spf13/cobra"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

const defaultOutputPrefix = "out"

// encodeCmd represents the encode command
var (
	imagePath  string
	outputPath string
	message    string
	encodeCmd  = &cobra.Command{
		Use:   "encode",
		Short: "encodes message in an image",
		Long: `
Stego encode will embed your message in an image.
It breaks the message into bits and writes them to
the least significant bit of each pixel's RGB channel
in the image.'`,
		RunE: encode,
	}
)

func writeImage(img image.Image, format string) error {
	if outputPath == "" {
		path, ext := utils.PathWithExtension(imagePath)
		outPath := strings.Join([]string{path, defaultOutputPrefix}, "-")
		outputPath = strings.Join([]string{outPath, ext}, ".")
	}
	writer, err := os.Create(outputPath)
	if writer != nil {
		defer writer.Close()
		defer writer.Sync()
	}
	if err != nil {
		return fmt.Errorf("error creating output image: %v", err)
	}
	if format == "png" {
		err = png.Encode(writer, img)
	} else if format == "jpeg" {
		err = jpeg.Encode(writer, img, &jpeg.Options{Quality: 100})
	}
	if err != nil {
		return fmt.Errorf("error encoding output image: %v", err)
	}
	return nil
}

func encode(cmd *cobra.Command, args []string) error {
	img, format, err := utils.GetImage(imagePath)
	if err != nil {
		return err
	}
	outImg, err := stego_lsb.LSBEncode(img, message)
	if err != nil {
		return fmt.Errorf("error encoding message: %v", err)
	}
	err = writeImage(outImg, format)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(encodeCmd)

	encodeCmd.Flags().StringVarP(&imagePath, "image", "i", "", "Path to the image")
	encodeCmd.MarkFlagRequired("image")
	encodeCmd.Flags().StringVarP(&message, "message", "m", "", "Message to be encoded")
	encodeCmd.MarkFlagRequired("message")
	encodeCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Output path for the encoded image")
}
