package main

import (
	"fmt"
	"os"
)

func pr(f string, v ...interface{}) {
	fmt.Printf(f+"\n", v...)
}

func check(e error) {
	if e != nil {
		pr("Error: %s", e.Error())
		os.Exit(2)
	}
}

func loadCurrentOrFail() *Actions {
	cwd, err := os.Getwd()
	check(err)
	fname, err := PathToFile(cwd)
	check(err)
	act, err := LoadActions(fname)
	check(err)
	return act
}
