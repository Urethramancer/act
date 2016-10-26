// +build linux dragonfly freebsd netbsd openbsd solaris

package main

import (
	"os"
	"path/filepath"
)

func configDir() string {
	path := filepath.Join(os.Getenv("HOME"), "."+program)
	return path
}
