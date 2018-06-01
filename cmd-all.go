package main

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/Urethramancer/cross"
)

// AllCmd tool command.
type AllCmd struct {
}

// Execute all.
func (ac *AllCmd) Execute(args []string) error {
	list, err := AllPaths()
	if err != nil {
		return err
	}

	for _, e := range list {
		path := filepath.Join(cross.ConfigPath(), e)
		act, err := LoadActions(path)
		if err != nil {
			return err
		}
		name := strings.TrimSuffix(e, ".json")
		pr("%s:", name)
		var a []string
		// We want all entries to have the colon aligned, so calculate a decent width.
		s := fmt.Sprintf("%d", act.Counter)
		w := len(s)
		width := fmt.Sprintf("%d", w)
		pr("Format: %s", width)
		for k, v := range act.List {
			s := fmt.Sprintf("%"+width+"d: %s", k, v)
			a = append(a, s)
		}
		sort.Strings(a)
		for _, x := range a {
			pr("%s", x)
		}
	}
	return nil
}
