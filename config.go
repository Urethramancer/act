package main

import (
	"os"
	"path/filepath"
)

const (
	program = "act"
	version = "0.1.1"
)

var _cfgdir, _dbname string

func defaultDB() string {
	if _dbname == "" || _cfgdir == "" {
		_cfgdir := configDir()
		_dbname = filepath.Join(_cfgdir, "action.db")
		if !exists(_cfgdir) {
			os.MkdirAll(_cfgdir, 0700)
		}
	}
	return _dbname
}
