package cmd

import (
	"fmt"

	"github.com/igorarthur/macleaner/internal/fs"
	"github.com/igorarthur/macleaner/internal/paths"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan for System Data trash in MacOS",
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, p := range paths.DockerPaths {
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
