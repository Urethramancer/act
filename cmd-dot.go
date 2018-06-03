package main

type DotCmd struct{}

func (dc *DotCmd) Execute(args []string) error {
	act := loadCurrentOrFail()
	if len(act.List) == 0 {
		pr("No actions for the current directory.")
		return nil
	}
	act.PrintActions(false, opts.Clean)
	return nil
}
