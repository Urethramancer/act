package main

type ClearCmd struct {
	Pos struct {
		Paths []string
	} `positional-args:"yes"`
}

func (cc *ClearCmd) Execute(args []string) error {
	return nil
}
