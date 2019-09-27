package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(inCmd)
}

var inCmd = &cobra.Command{
	Use:   "in",
	Short: "clock in",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: implement function for clocking in
		fmt.Println("in!")
	},
}
