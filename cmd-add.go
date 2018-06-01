package main

type AddCmd struct {
	Action string
}

func (ac *AddCmd) Execute(args []string) error {
	return nil
}
