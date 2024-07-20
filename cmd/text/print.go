package text

import (
	"fmt"
	"github.com/brlywk/termup/cmd/styles"
	"github.com/charmbracelet/lipgloss"
)

// ColorLn is a wrapper around fmt.Printf, using styles.RenderColorString to
// text colored text to stdout.
func ColorLn(color lipgloss.Color, s string, a ...any) {
	fmt.Println(styles.RenderColorString(color, fmt.Sprintf(s, a...)))
}
