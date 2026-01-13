package cmd

import (
	"fmt"
	"runtime"

	"github.com/igorarthur/macleaner/internal/fs"
	"github.com/igorarthur/macleaner/internal/paths"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan for System Data trash in your system",
	RunE: func(cmd *cobra.Command, args []string) error {
		goos := runtime.GOOS
		DockerPaths := paths.DockerPaths[goos]

		for _, p := range DockerPaths {
			expanded, err := fs.ExpandPath(p)
			if err != nil {
				continue
			}

			size, err := fs.DirSize(expanded)
			if err != nil {
				continue
			}

			fmt.Printf("%s → %.2f GB\n", expanded, float64(size)/1024/1024/1024)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
