// act
// A TODO list for the command line.

package main

import (
	"os"

	"github.com/Urethramancer/cross"
	"github.com/jessevdk/go-flags"
)

func main() {
	cross.SetConfigPath(program)
	if !cross.DirExists(cross.ConfigPath()) {
		pr("Couldn't create folder '%s'", cross.ConfigPath())
		os.Exit(2)
	}

	if len(os.Args) == 1 {
		act := loadCurrentOrFail()
		if len(act.List) == 0 {
			pr("No actions for the current directory.")
			return
		}
		act.PrintActions(false, false)
		return
	}

	if len(os.Args) == 2 && os.Args[1] == "-c" {
		act := loadCurrentOrFail()
		if len(act.List) == 0 {
			pr("No actions for the current directory.")
			return
		}
		act.PrintActions(false, true)
		return
	}

	_, _ = flags.Parse(&opts)
}
