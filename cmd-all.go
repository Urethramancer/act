package main

import (
	"path/filepath"

	"github.com/Urethramancer/cross"
)

// AllCmd tool command.
type AllCmd struct{}

// Execute all.
func (ac *AllCmd) Execute(args []string) error {
	list, err := AllPaths()
	if err != nil {
		return err
	}

	for _, e := range list {
		path := filepath.Join(cross.ConfigPath(), e)
		act, err := LoadActions(path)
		if err != nil {
			return err
		}
		act.PrintActions(true, opts.Clean)
	}
	return nil
}
