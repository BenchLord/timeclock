package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(outCmd)
}

var outCmd = &cobra.Command{
	Use:   "out",
	Short: "clock out",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: implement function for clocking out
		fmt.Println("out!")
	},
}
