package main

import (
	"io/ioutil"
	"regexp"

	"github.com/Urethramancer/cross"
)

func PathToFile(path string) (string, error) {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return "", err
	}

	res := reg.ReplaceAllString(path, "-")
	res = res[1:]
	res = cross.ConfigName(res)
	return res + ".json", nil
}

func AllPaths() ([]string, error) {
	var list []string
	path := cross.ConfigPath()
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, e := range entries {
		if !e.IsDir() {
			list = append(list, e.Name())
		}
	}
	return list, nil
}
