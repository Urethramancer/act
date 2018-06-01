package main

// ClearCmd tool command.
type ClearCmd struct {
	Pos struct {
		Paths []string
	} `positional-args:"yes"`
}

// Execute clear.
func (cc *ClearCmd) Execute(args []string) error {
	return nil
}
