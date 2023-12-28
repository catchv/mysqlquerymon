package utils

import (
	"fmt"
	"os"
)

type Color string

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed    Color = "\u001b[31m"
	ColorGreen  Color = "\u001b[32m"
	ColorYellow Color = "\u001b[33m"
	ColorBlue   Color = "\u001b[34m"
	ColorPurple Color = "\u001b[35m"
	ColorReset  Color = "\u001b[0m"
)

func ColorPrintln(color Color, message string) {
	fmt.Fprintf(os.Stdout, "%s%s%s\n", color, message, ColorReset)
}
