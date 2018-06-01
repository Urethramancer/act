package main

// RemoveCmd tool command.
type RemoveCmd struct {
	Pos struct {
		IDs []string
	} `positional-args:"yes"`
}

// Execute remove.
func (rc *RemoveCmd) Execute(args []string) error {
	return nil
}
