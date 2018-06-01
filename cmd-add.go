package main

// AddCmd tool command.
type AddCmd struct {
	Action string
}

// Execute add.
func (ac *AddCmd) Execute(args []string) error {
	return nil
}
