package program

type Homebrew struct {
	defaultProgram
}

func (h *Homebrew) available() bool {
	return h.defaultProgram.available()
}

func (h *Homebrew) init() {
	h.name = "homebrew"
	h.command = "brew"
	h.installScript = `/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"`
}

func (h *Homebrew) Install() error {
	h.init()

	return h.defaultProgram.Install()
}
