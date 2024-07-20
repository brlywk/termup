package program

import (
	"os"
	"os/exec"
)

const (
	shell     = "/bin/zsh"
	shellFlag = "-c"
)

// commandAvailable checks if a specific command is found in PATH.
func commandAvailable(cmdName string) bool {
	_, err := exec.LookPath(cmdName)
	return err == nil
}

func runCmd(command string) error {
	cmd := exec.Command(shell, shellFlag, command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// InstallPrerequisites checks for and installs required programs.
func InstallPrerequisites() error {
	var homebrew Homebrew

	err := homebrew.Install()
	if err != nil {
		return err
	}

	var chezmoi Chezmoi
	err = chezmoi.Install()
	if err != nil {
		return err
	}

	return nil
}
