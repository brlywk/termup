package program

import (
	"github.com/brlywk/termup/cmd/styles"
	"github.com/brlywk/termup/cmd/text"
)

type Homebrew struct {
	command       string
	installScript string
}

func (h *Homebrew) available() bool {
	return commandAvailable(h.command)
}

func (h *Homebrew) init() {
	h.command = "brew"
	h.installScript = `/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"`
}

func (h *Homebrew) Install() error {
	h.init()

	if h.available() {
		text.ColorLn(styles.BrightBlack, "Homebrew already installed.")
		return nil
	}

	return runCmd(h.installScript)
}
