package cmd

import (
	"fmt"
	"os"

	"github.com/igorarthur/macleaner/internal/fs"
	"github.com/igorarthur/macleaner/internal/paths"
	"github.com/spf13/cobra"
)

var dryRun bool
var assumeYes bool

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Remove Docker-related system files",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !assumeYes {
			fmt.Println("This will permanently delete Docker files.")
			fmt.Print("Continue? [y/N]: ")

			var confirm string
			fmt.Scanln(&confirm)
			if confirm != "y" && confirm != "Y" {
				return nil
			}
		}

		for _, p := range paths.DockerPaths {
			expanded, err := fs.ExpandPath(p)
			if err != nil {
				continue
			}

			if dryRun {
				fmt.Println("[dry-run]", expanded)
				continue
			}

			err = os.RemoveAll(expanded)
			if err != nil {
				fmt.Println("Failed:", expanded, err)
			} else {
				fmt.Println("Removed:", expanded)
			}
		}

		return nil
	},
}

func init() {
	cleanCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Show what would be deleted")
	cleanCmd.Flags().BoolVarP(&assumeYes, "yes", "y", false, "Skip confirmation")

	rootCmd.AddCommand(cleanCmd)
}
