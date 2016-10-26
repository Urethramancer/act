package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"

	"github.com/asdine/storm"
)

func editAction(db *storm.DB, action string) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		pr("No editor configured. Set your EDITOR environment variable.")
		return
	}

	var act Action
	n, err := strconv.Atoi(action)
	check(err)
	err = db.One("ID", n, &act)
	check(err)

	f, err := ioutil.TempFile(os.TempDir(), "act")
	check(err)
	f.WriteString(act.Text)
	fn := f.Name()
	f.Close()

	cmd := exec.Command(editor, fn)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	check(err)
	err = cmd.Wait()
	check(err)

	text, err := ioutil.ReadFile(fn)
	os.Remove(fn)
	check(err)

	act.Text = string(text)
	err = db.Update(&act)
	check(err)
}
