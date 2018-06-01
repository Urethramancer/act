package main

import (
	"io/ioutil"
	"os"
	"os/exec"
)

type EditCmd struct {
	Action string
	Text   string
	Editor bool `short:"e" long:"editor" description:"Use an editor instead of text entered as an argument."`
}

func (ec *EditCmd) Execute(args []string) error {
	return nil
}

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
