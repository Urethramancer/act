package main

const (
	program = "act"
)

// Version is filled in from Git tags.
var Version = "undefined"

// VersionCmd just holds the Execute() command.
type VersionCmd struct{}

// Execute the version command.
func (vc *VersionCmd) Execute(args []string) error {
	pr("%s %s", program, Version)
	return nil
}
