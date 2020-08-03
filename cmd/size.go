package cmd

import (
	"fmt"
	"github.com/gzcharleszhang/stego/pkg/stegolsb"
	"github.com/gzcharleszhang/stego/utils"
	"math"

	"github.com/spf13/cobra"
)

// sizeCmd represents the size command
var (
	sizePrettyPrint bool
	sizeCmd         = &cobra.Command{
		Use:   "size",
		Short: "Maximum encoding size of an image",
		Long: `
This command calculates the maximum amount of data
the given image is able to encode.`,
		Args: cobra.ExactArgs(1),
		RunE: runSize,
	}
)

// used for pretty printing bytes
func prettyBytes(bytes uint32) string {
	units := []string{"KB", "MB", "GB", "TB"}
	curr := float64(bytes)
	currUnit := "B"
	for _, unit := range units {
		next := curr / 1024
		if next < 1 {
			break
		}
		curr = next
		currUnit = unit
	}
	result := math.Round(curr*100) / 100
	return fmt.Sprintf("%v %v", result, currUnit)
}

func runSize(cmd *cobra.Command, args []string) error {
	imagePath := args[0]
	img, _, err := utils.GetImage(imagePath)
	if err != nil {
		return err
	}
	size, err := stegolsb.MaxLSBEncodeSize(img)
	if sizePrettyPrint {
		cmd.Println(prettyBytes(size))
	} else {
		cmd.Println(size)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(sizeCmd)
	sizeCmd.Flags().BoolVarP(&sizePrettyPrint, "pretty", "p", false, "Pretty print the maximum encoding size")
}
