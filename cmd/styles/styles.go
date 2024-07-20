package styles

import "github.com/charmbracelet/lipgloss"

var (
	HeaderStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(BrightCyan).
		//Background(Black).
		Padding(1, 2)
)
