package main

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/Urethramancer/cross"
)

type Actions struct {
	Counter      uint64            `json:"counter"`
	LastModified time.Time         `json:"modified"`
	List         map[string]string `json:"actions"`
}

func LoadActions(name string) (*Actions, error) {
	var err error
	if !cross.FileExists(name) {
		act := Actions{
			Counter:      1,
			LastModified: time.Now(),
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

func SaveActions(name string, act *Actions) error {
	data, err := json.MarshalIndent(act, "", "\t")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(name, data, 0600)
}