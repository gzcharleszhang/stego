package cmd

import (
	"fmt"
	"github.com/gzcharleszhang/stego/pkg/stegolsb"
	"github.com/gzcharleszhang/stego/utils"

	"github.com/spf13/cobra"
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "decodes data in an image",
	Long: `
Stego decode attempts to retrieve the hidden data in an image.
Decode looks at the least significant bit of each of the RGB channels
in each pixel of the image, and collects those bits to recompose
the data. Decode treats the first 8 bits of the data as the total size
of the hidden data. Decode looks at more significant bits after
going through the least significant bits.`,
	Args: cobra.ExactArgs(1),
	RunE: run,
}

func run(cmd *cobra.Command, args []string) error {
	imagePath := args[0]
	img, _, err := utils.GetImage(imagePath)
	if err != nil {
		return err
	}
	message, err := stegolsb.Decode(img)
	if err != nil {
		return fmt.Errorf("error decoding image: %v", err)
	}
	cmd.Println(message)
	return nil
}

func init() {
	rootCmd.AddCommand(decodeCmd)
}
