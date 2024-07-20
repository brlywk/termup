package program

type Program interface {
	// check checks if the command is available.
	available() bool
	// init is used to set up the struct internally.
	init()
	// Install checks if the command is available and installs it if necessary.
	Install() error
}
