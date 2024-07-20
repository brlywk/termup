package cmd

import (
	"fmt"
	"github.com/brlywk/termup/cmd/program"
	"github.com/brlywk/termup/cmd/styles"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "termup",
	Short: "termup - the TERMinal setUP tool.",
	Long:  `termup helps setting up your terminal, with the help of homebrew and chezmoi.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := program.InstallPrerequisites()
		if err != nil {
			cmd.PrintErrln(styles.RenderColorString(styles.Red, fmt.Sprintf("Error installing prerequisite: %v", err)))
			os.Exit(1)
		}
	},
}

// Execute runs the cobra root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
