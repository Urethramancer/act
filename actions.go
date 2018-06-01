package main

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/Urethramancer/cross"
)

// Actions structure holds the actions for one directory.
type Actions struct {
	// Counter is the number of the next entry to be added.
	Counter uint64 `json:"counter"`
	// LastModified is in local time.
	LastModified time.Time `json:"modified"`
	// List contains the actual actions.
	List map[string]string `json:"actions"`
}

// LoadActions loads a JSON file representing an Actions structure.
func LoadActions(name string) (*Actions, error) {
	var err error
	if !cross.FileExists(name) {
		act := Actions{
			Counter:      1,
			LastModified: time.Now().Local(),
			List:         make(map[string]string),
		}
		err = SaveActions(name, &act)
		return &act, err
	}

	data, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}

	var act Actions
	err = json.Unmarshal(data, &act)
	return &act, err
}

// SaveActions saves a nicely formatted version of an Actions structure.
func SaveActions(name string, act *Actions) error {
	data, err := json.MarshalIndent(act, "", "\t")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(name, data, 0600)
}
