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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
