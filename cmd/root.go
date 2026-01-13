package cmd

import (
	"os"

	"github.com/igorarthur/macleaner/internal/support"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "macleaner",
	Short: "MaCleaner",
	Long:  "MaCleaner is a command-line tool for cleaning up Docker-related files on macOS, Linux and Windows.",
}

func Execute() {
	support.EnsureSupportedOS()

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
