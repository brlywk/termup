package program

type Git struct {
	defaultProgram
}

// check checks if the command is available.
func (g *Git) available() bool {
	return g.defaultProgram.available()
}

// init is used to set up the struct internally.
func (g *Git) init() {
	g.defaultProgram.initDefault(
		"git",
		"git",
		"brew install git",
		&Homebrew{},
	)
}

// Install checks if the command is available and installs it if necessary.
func (g *Git) Install() error {
	g.init()

	return g.defaultProgram.Install()
}
