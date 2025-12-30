package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Diagnose Docker-related disk usage on macOS",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print("MaCleaner Doctor\n")

		checkDockerRunning()
		explainSystemData()
		explainLimitations()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}

func checkDockerRunning() {
	cmd := exec.Command("pgrep", "-f", "Docker Desktop")
	if err := cmd.Run(); err == nil {
		fmt.Println("✔ Docker Desktop is running")
	} else {
		fmt.Println("ℹ Docker Desktop is not running")
	}
}

func explainSystemData() {
	fmt.Println("\nWhy 'System Data' grows with Docker:")
	fmt.Println("- Docker stores container layers and volumes in a hidden Linux VM")
	fmt.Println("- macOS reports this as 'System Data'")
	fmt.Println("- Deleting Finder-visible files does not affect this usage")
}

func explainLimitations() {
	fmt.Println("\nWhat this tool CAN clean:")
	fmt.Println("- Docker Desktop caches")
	fmt.Println("- Logs and application support files")
	fmt.Println("- Group containers left after image rebuilds")

	fmt.Println("\nWhat this tool CANNOT clean:")
	fmt.Println("- Docker VM disk images")
	fmt.Println("- APFS snapshots")
	fmt.Println("- Files protected by SIP")

	fmt.Println("\nRecommended actions:")
	fmt.Println("- Run: docker system prune -a")
	fmt.Println("- Restart Docker Desktop")
	fmt.Println("- Quit Docker when not in use")
}
