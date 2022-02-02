package printer

import "strings"

type Color string

const (
	BLACK        Color = "\033[0;30m"
	RED          Color = "\033[0;31m"
	GREEN        Color = "\033[0;32m"
	BROWN        Color = "\033[0;33m"
	BLUE         Color = "\033[0;34m"
	PURPLE       Color = "\033[0;35m"
	CYAN         Color = "\033[0;36m"
	LIGHT_GRAY   Color = "\033[0;37m"
	DARK_GRAY    Color = "\033[1;30m"
	LIGHT_RED    Color = "\033[1;31m"
	LIGHT_GREEN  Color = "\033[1;32m"
	YELLOW       Color = "\033[1;33m"
	LIGHT_BLUE   Color = "\033[1;34m"
	LIGHT_PURPLE Color = "\033[1;35m"
	LIGHT_CYAN   Color = "\033[1;36m"
	LIGHT_WHITE  Color = "\033[1;37m"
	BOLD         Color = "\033[1m"
	FAINT        Color = "\033[2m"
	ITALIC       Color = "\033[3m"
	UNDERLINE    Color = "\033[4m"
	BLINK        Color = "\033[5m"
	NEGATIVE     Color = "\033[7m"
	CROSSED      Color = "\033[9m"
	END          Color = "\033[0m"
)

func Colorize(s string, c Color) string {
	// "ColorCode CurrentString EndCode"
	return string(c) + s + string(END)
}

func ColorizeWord(s string, word string, c Color) string {
	// "ColorCode word EndCode"
	replacer := string(c) + word + string(END)

	replaced := strings.ReplaceAll(s, word, replacer)

	return replaced
}

var ErrorMsg = Colorize("Sorry, couldn't find a valid quote!", RED)
var ThanksMsg = Colorize("Thank you for contributing", GREEN)
