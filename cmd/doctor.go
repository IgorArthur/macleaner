package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Diagnose Docker-related disk usage on your system",
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
	fmt.Println("\nWhy Docker storage grows across different OSs:")

	fmt.Println("\n[macOS]")
	fmt.Println("- Docker Desktop runs a hidden Linux VM (Docker.raw).")
	fmt.Println("- macOS 'System Data' includes this massive VM file.")
	fmt.Println("- Deleting files in Finder won't shrink the VM disk image automatically.")

	fmt.Println("\n[Windows]")
	fmt.Println("- Docker typically uses WSL2 (Windows Subsystem for Linux).")
	fmt.Println("- Data is stored in a virtual disk file (ext4.vhdx).")
	fmt.Println("- This shows up as 'Other' or 'System' storage in Windows settings.")

	fmt.Println("\n[Linux]")
	fmt.Println("- Docker runs natively; there is no hidden VM.")
	fmt.Println("- Data stays in /var/lib/docker.")
	fmt.Println("- While it's not 'System Data,' it can fill the root partition (/) quickly.")

	fmt.Println("\nNote: Use 'docker system prune' on any OS to reclaim this space.")
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
