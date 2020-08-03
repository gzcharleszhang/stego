/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	stego_lsb "github.com/gzcharleszhang/stego/pkg/stego-lsb"
	"github.com/gzcharleszhang/stego/utils"
	"github.com/spf13/cobra"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

const DEFAULT_OUTPUT_PREFIX = "out"

// encodeCmd represents the encode command
var (
	imagePath string
	outputPath string
	message string
	encodeCmd = &cobra.Command{
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
		outPath := strings.Join([]string{path, DEFAULT_OUTPUT_PREFIX}, "-")
		outputPath = strings.Join([]string{outPath, ext}, ".")
	}
	writer, err := os.Create(outputPath)
	if writer != nil {
		defer writer.Close()
		defer writer.Sync()
	}
	if err != nil {
		return fmt.Errorf("Error creating output image: %v\n", err)
	}
	if format == "png" {
		err = png.Encode(writer, img)
	} else if format == "jpeg" {
		err = jpeg.Encode(writer, img, &jpeg.Options{Quality: 100})
	}
	if err != nil {
		return fmt.Errorf("Error encoding output image: %v\n", err)
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
		return fmt.Errorf("Error encoding message: %v\n", err)
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
