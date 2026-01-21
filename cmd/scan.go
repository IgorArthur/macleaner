package cmd

import (
	"fmt"
	"io"
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

		return Scan(realFS, os.Stdout)
	},
}

func Scan(fs fs.FileSystem, out io.Writer) error {
	goos := runtime.GOOS
	DockerPaths := paths.DockerPaths[goos]
	var found int

	for _, p := range DockerPaths {
		expanded, err := fs.ExpandPath(p)
		if err != nil {
			continue
		}

		if !fs.Exists(expanded) {
			continue
		}
		found++

		size, err := fs.DirSize(expanded)
		if err != nil {
			continue
		}

		fmt.Fprintf(out, "%s → %.2f GB\n", expanded, float64(size)/1024/1024/1024)
	}

	if found == 0 {
		fmt.Fprintf(out, "No Docker paths found in your %s system\n", goos)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
