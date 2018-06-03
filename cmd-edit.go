package main

import (
	"io/ioutil"
	"os"
	"os/exec"
)

// EditCmd tool command.
type EditCmd struct {
	Args struct {
		IDs []uint64 `positional-arg-name:"ACTION" description:"ID of action to edit."`
	} `positional-args:"yes"`
}

// Execute edit.
func (ec *EditCmd) Execute(args []string) error {
	if len(ec.Args.IDs) == 0 {
		pr("You need to specify one or more action IDs to edit.")
		return nil
	}

	act := loadCurrentOrFail()
	pr("%s", act.filename)
	return nil
}

// EditText opens the supplied string with the text editor in the EDITOR environment variable.
func EditText(in string) string {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		pr("No editor configured. Set your EDITOR environment variable.")
		return in
	}

	f, err := ioutil.TempFile(os.TempDir(), "act")
	check(err)
	_, _ = f.WriteString(in)
	fn := f.Name()
	err = f.Close()
	if err != nil {
		return in
	}

	cmd := exec.Command(editor, fn)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		return in
	}

	err = cmd.Wait()
	if err != nil {
		return in
	}

	out, err := ioutil.ReadFile(fn)
	_ = os.Remove(fn)
	if err != nil {
		return in
	}

	return string(out)
}
