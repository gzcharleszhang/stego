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
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: run,
}

func run(cmd *cobra.Command, args []string) error {
	img, _, err := utils.GetImage(imagePath)
	if err != nil {
		return err
	}
	message, err := stego_lsb.Decode(img)
	if err != nil {
		return fmt.Errorf("Error decoding image: %v\n", err)
	}
	cmd.Println(message)
	return nil
}

func init() {
	rootCmd.AddCommand(decodeCmd)

	decodeCmd.Flags().StringVarP(&imagePath, "image", "i", "", "Path to the image")
	decodeCmd.MarkFlagRequired("image")
}
