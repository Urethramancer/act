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

	_, _ = flags.Parse(&opts)
}
