package main

import (
	"os"
	"path/filepath"
)

func configDir() string {
	path := filepath.Join(os.Getenv("HOME"), "Library", "Application Support", program)
	return path
}
