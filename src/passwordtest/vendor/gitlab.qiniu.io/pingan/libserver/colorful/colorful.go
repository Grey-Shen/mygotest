package colorful

import (
	"os"

	isatty "github.com/mattn/go-isatty"

	"github.com/labstack/gommon/color"
)

var (
	StdoutColor *color.Color
	StderrColor *color.Color
)

func init() {
	StdoutColor = color.New()
	StderrColor = color.New()

	if isatty.IsTerminal(os.Stdout.Fd()) {
		StdoutColor.Enable()
	} else {
		StdoutColor.Disable()
	}

	if isatty.IsTerminal(os.Stderr.Fd()) {
		StderrColor.Enable()
	} else {
		StderrColor.Disable()
	}
}
