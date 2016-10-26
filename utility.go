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

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
