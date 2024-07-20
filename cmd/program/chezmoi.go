package program

type Chezmoi struct {
	defaultProgram
}

func (c *Chezmoi) available() bool {
	return c.defaultProgram.available()
}

func (c *Chezmoi) init() {
	c.defaultProgram.initDefault(
		"chezmoi",
		"chezmoi",
		"brew install chezmoi",
		&Homebrew{},
	)
}

func (c *Chezmoi) Install() error {
	c.init()

	return c.defaultProgram.Install()
}
