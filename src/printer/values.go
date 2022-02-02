package printer

import (
	"fmt"

	"github.com/force4760/quotedopai/version"
)

func ColorVersion() {
	fmt.Println(Colorize("Build Date:", BLUE), version.BuildDate)
	fmt.Println(Colorize("Git Commit:", BLUE), version.GitCommit)
	fmt.Println(Colorize("Version:", BLUE), version.Version)
	fmt.Println(Colorize("Go Version:", BLUE), version.GoVersion)
	fmt.Println(Colorize("OS / Arch:", BLUE), version.OsArch)
}

var ErrorMsg = Colorize("Sorry, couldn't find a valid quote!", RED)
var ThanksMsg = Colorize("Thank you for contributing", GREEN)
