package cmd

import (
	"fmt"
	"os"
	"runtime"

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
		realFS := &fs.RealFS{}

		return Clean(realFS, dryRun, assumeYes)
	},
}

func Clean(fs fs.FileSystem, dryRun bool, assumeYes bool) error {
	if !assumeYes {
		fmt.Println("This will permanently delete Docker files.")
		fmt.Print("Continue? [y/N]: ")

		var confirm string
		fmt.Scanln(&confirm)
		if confirm != "y" && confirm != "Y" {
			return nil
		}
	}

	goos := runtime.GOOS
	DockerPaths := paths.DockerPaths[goos]
	var found int

	for _, p := range DockerPaths {
		expanded, err := fs.ExpandPath(p)
		if err != nil {
		}

		if !fs.Exists(expanded) {
			continue
		}

		if dryRun {
			found++
			fmt.Fprintln(os.Stdout, "[dry-run]", expanded)
			continue
		}

		err = fs.RemoveAll(expanded)
		if err != nil {
			fmt.Fprintln(os.Stdout, "Failed:", expanded, err)
		} else {
			found++
			fmt.Fprintln(os.Stdout, "Removed:", expanded)
		}
	}

	if found == 0 {
		fmt.Fprintln(os.Stdout, "No Docker paths found in your %s system\n", goos)
	}

	return nil
}

func init() {
	cleanCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Show what would be deleted")
	cleanCmd.Flags().BoolVarP(&assumeYes, "yes", "y", false, "Skip confirmation")

	rootCmd.AddCommand(cleanCmd)
}
