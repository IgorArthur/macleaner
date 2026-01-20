package support

import (
	"fmt"
	"os"
	"runtime"
	"slices"
)

func EnsureSupportedOS() {
	supported := []string{"darwin", "windows", "linux"}

	if !slices.Contains(supported, runtime.GOOS) {
		fmt.Printf("❌ Fatal: Operating system '%s' is not supported by this tool.\n", runtime.GOOS)
		os.Exit(1)
	}
}
