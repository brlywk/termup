package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setupCmd)
}

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup environment",
	Long: `Setup runs the full setup for an environment.
It will check that both homebrew and chezmoi are available,
and installs them if necessary.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("setup called")
	},
}
