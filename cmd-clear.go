package main

import (
	"errors"
	"os"

	"github.com/Urethramancer/cross"
)

// ClearCmd tool command.
type ClearCmd struct {
	Args struct {
		Path string `positional-arg-name:"PATH" description:"Path to clear actions for."`
	} `positional-args:"yes"`
}

// Execute clear.
func (cc *ClearCmd) Execute(args []string) error {
	if cc.Args.Path == "" {
		pr("You need to specify a path to clear. You can clear the current path with \".\"")
		return nil
	}

	var fname string
	var err error
	if cc.Args.Path == "." {
		cwd, err := os.Getwd()
		check(err)
		fname, err = PathToFile(cwd)
	} else {
		fname, err = PathToFile(cc.Args.Path)
	}
	check(err)

	if !cross.FileExists(fname) {
		return errors.New("no such folder")
	}

	return os.Remove(fname)
}
