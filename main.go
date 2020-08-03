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
package main

import (
	"fmt"
	"github.com/gzcharleszhang/stego/pkg/stego-lsb"
	"image"
	"image/png"
	"os"
)

func example() {
	reader, err := os.Open("./test/testdata/butterfly.png")
	if reader != nil {
		defer reader.Close()
	}
	if err != nil {
		fmt.Printf("Error opening png file: %v\n", err)
	}
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	img, _, err := image.Decode(reader)
	if err != nil {
		fmt.Printf("Error decoding image: %v\n", err)
	}
	maxSize, _ := stego_lsb.MaxLSBEncodeSize(img)
	fmt.Printf("Maximum encode size: %d bytes\n", maxSize)
	outImg, err := stego_lsb.LSBEncode(img, "Hello, world!")
	if err != nil {
		fmt.Printf("Error encoding message: %v\n", err)
	}
	writer, err := os.Create("./test/testdata/butterfly-out.png")
	if writer != nil {
		defer writer.Close()
		defer writer.Sync()
	}
	if err != nil {
		fmt.Printf("Error creating png file: %v\n", err)
	}
	err = png.Encode(writer, outImg)
	if err != nil {
		fmt.Printf("Error encoding png image to file: %v\n", err)
	}
	message, err := stego_lsb.Decode(outImg)
	if err != nil {
		fmt.Printf("Error decoding image: %v\n", err)
	}
	fmt.Println(message)
}

func main() {
	example()
}
