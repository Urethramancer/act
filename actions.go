package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/Urethramancer/cross"
)

// Actions structure holds the actions for one directory.
type Actions struct {
	filename string
	// Counter is the number of the most recent entry.
	Counter uint64 `json:"counter"`
	// LastModified is in local time.
	LastModified time.Time `json:"modified"`
	// List contains the actual actions.
	List map[uint64]string `json:"actions"`
}

// LoadActions loads a JSON file representing an Actions structure.
func LoadActions(name string) (*Actions, error) {
	var err error
	if !cross.FileExists(name) {
		act := Actions{
			filename:     name,
			Counter:      0,
			LastModified: time.Now().Local(),
			List:         make(map[uint64]string),
		}
		err = act.Save()
		return &act, err
	}

	data, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}

	var act Actions
	err = json.Unmarshal(data, &act)
	act.filename = name
	return &act, err
}

// SaveActions saves a nicely formatted version of an Actions structure.
func (act *Actions) Save() error {
	data, err := json.MarshalIndent(act, "", "\t")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(act.filename, data, 0600)
}

func (act *Actions) PrintActions(indent bool) {
	if indent {
		name := filepath.Base(act.filename)
		name = strings.TrimSuffix(name, ".json")
		pr("%s:", name)
	}

	var a []string
	// We want all entries to have the colon aligned, so calculate a decent width.
	s := fmt.Sprintf("%d", act.Counter)
	w := len(s)
	width := fmt.Sprintf("%d", w)
	for k, v := range act.List {
		s := fmt.Sprintf("%"+width+"d: %s", k, v)
		a = append(a, s)
	}
	sort.Strings(a)
	for _, x := range a {
		if indent {
			fmt.Print("\t")
		}
		pr("%s", x)
	}
}
