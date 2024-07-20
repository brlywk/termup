package program

import (
	"os"
	"os/exec"
)

const (
	shell     = "/bin/zsh"
	shellFlag = "-c"
)

// prerequisites to install
var prerequisites = []Program{
	&Homebrew{},
	&Git{},
	&Chezmoi{},
}

// commandAvailable checks if a specific command is found in PATH.
func commandAvailable(cmdName string) bool {
	_, err := exec.LookPath(cmdName)
	return err == nil
}

// runCmd runs command (using package shell and shellFlags).
func runCmd(command string) error {
	cmd := exec.Command(shell, shellFlag, command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// InstallPrerequisites checks for and installs required programs.
func InstallPrerequisites() error {
	for _, prereq := range prerequisites {
		err := prereq.Install()
		if err != nil {
			return err
		}
	}

	return nil
}
