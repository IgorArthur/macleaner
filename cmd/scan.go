package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/igorarthur/macleaner/internal/fs"
	"github.com/igorarthur/macleaner/internal/paths"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan for System Data trash in your system",
	RunE: func(cmd *cobra.Command, args []string) error {
		realFS := &fs.RealFS{}

		return scan(realFS)
	},
}

func scan(fs fs.FileSystem) error {
	goos := runtime.GOOS
	DockerPaths := paths.DockerPaths[goos]
	var found int

	for _, p := range DockerPaths {
		expanded, err := fs.ExpandPath(p)
		if err != nil {
			continue
		}

		if _, err := os.Stat(expanded); os.IsNotExist(err) {
			continue
		}
		found++

		size, err := fs.DirSize(expanded)
		if err != nil {
			continue
		}

		fmt.Printf("%s → %.2f GB\n", expanded, float64(size)/1024/1024/1024)
	}

	if found == 0 {
		fmt.Printf("No Docker paths found in your %s system\n", goos)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
