package cmd

import (
	"github.com/AndresRamirez9912/pattern-of-the-day/internal"
	"github.com/spf13/cobra"
)

// NewVersionInfoCmd creates a command to display version information
func NewVersionInfoCmd() *cobra.Command {
	// Create the version info command
	versionInfoCmd := &cobra.Command{
		Use:   "version",
		Short: "Display version information",
		Long:  `Display the current version, commit hash, and build date of the KiiEx backend.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("Version:", internal.Version)
			cmd.Println("Commit:", internal.Commit)
			cmd.Println("Build Date:", internal.BuildDate)
		},
	}

	// Return the version info command
	return versionInfoCmd
}
