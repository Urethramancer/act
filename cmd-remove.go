package main

type RemoveCmd struct {
	Pos struct {
		IDs []string
	} `positional-args:"yes"`
}

func (rc *RemoveCmd) Execute(args []string) error {
	return nil
}
