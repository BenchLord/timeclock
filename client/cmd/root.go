package cmd

import (
	"fmt"
	"os"

	dialer "github.com/benchlord/timeclock/client/internal"
	"github.com/spf13/cobra"
)

var (
	client *dialer.Client
	err    error
)

var rootCmd = &cobra.Command{
	Use: "timeclock",
}

// Execute runs the root command and exits on error.
func Execute() {
	client, err = dialer.NewClient()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
