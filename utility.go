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
