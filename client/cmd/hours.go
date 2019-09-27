package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(hoursCmd)
}

var hoursCmd = &cobra.Command{
	Use:   "hours",
	Short: "get total hours worked",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: implement function for getting hours
		fmt.Println("hours!")
	},
}
