package program

import (
	"github.com/brlywk/termup/cmd/styles"
	"github.com/brlywk/termup/cmd/text"
)

type Program interface {
	// check checks if the command is available.
	available() bool
	// init is used to set up the struct internally.
	init()
	// Install checks if the command is available and installs it if necessary.
	Install() error
}

// defaultProgram is a base implementation for Program used by most structs implementing Program.
type defaultProgram struct {
	name          string
	requires      []Program
	command       string
	installScript string
}

// check checks if the command is available.
func (d *defaultProgram) available() bool {
	return commandAvailable(d.command)
}

// initDefault is an overwrite for init() used only for the defaultProgram, as most
// Programs defined here have pretty much the same available, init and Install methods.
func (d *defaultProgram) initDefault(name, command, installScript string, requires ...Program) {
	d.name = name
	d.command = command
	d.installScript = installScript
	d.requires = requires
}

// init is used to set up the struct internally.
func (d *defaultProgram) init() {}

// Install checks if the command is available and installs it if necessary.
func (d *defaultProgram) Install() error {
	if d.available() {
		text.ColorLn(styles.BrightBlack, "%s already installed.", d.name)

		return nil
	}

	for _, req := range d.requires {
		if !req.available() {
			err := req.Install()
			if err != nil {
				return err
			}
		}
	}

	return runCmd(d.installScript)
}
