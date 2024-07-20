package program

import (
	"github.com/brlywk/termup/cmd/styles"
	"github.com/brlywk/termup/cmd/text"
)

type Chezmoi struct {
	requires      []Program
	command       string
	installScript string
}

func (c *Chezmoi) available() bool {
	return commandAvailable(c.command)
}

func (c *Chezmoi) init() {
	c.command = "chezmoi"
	c.installScript = "brew install chezmoi"

	c.requires = []Program{
		&Homebrew{},
	}
}

func (c *Chezmoi) Install() error {
	c.init()

	if c.available() {
		text.ColorLn(styles.BrightBlack, "Chezmoi already installed.")

		return nil
	}

	for _, cmd := range c.requires {
		if !cmd.available() {
			err := cmd.Install()
			if err != nil {
				return err
			}
		}
	}

	return runCmd(c.installScript)
}
