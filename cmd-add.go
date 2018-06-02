package main

import (
	"os"
	"time"
)

// AddCmd tool command.
type AddCmd struct {
	Action struct {
		Text string `positional-arg-name:"TEXT" description:"Action text in quotes."`
	} `positional-args:"yes"`
}

// Execute add.
func (ac *AddCmd) Execute(args []string) error {
	if len(ac.Action.Text) == 0 {
		pr("You need to specify an action to add.")
		os.Exit(1)
	}

	act := loadCurrentOrFail()
	act.Counter++
	act.List[act.Counter] = ac.Action.Text
	act.LastModified = time.Now().Local()
	act.Save()
	return nil
}
