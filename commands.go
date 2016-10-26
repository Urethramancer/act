package main

import "strings"

type VersionCmd struct {
	Run bool
}

func (vc *VersionCmd) Execute(args []string) error {
	vc.Run = true
	return nil
}

type AllCmd struct {
	Run bool
}

func (ac *AllCmd) Execute(args []string) error {
	ac.Run = true
	return nil
}

type AddCmd struct {
	Action string
	Run    bool
}

func (ac *AddCmd) Execute(args []string) error {
	ac.Run = true
	ac.Action = strings.Join(args, " ")
	return nil
}

type RemoveCmd struct {
	Pos struct {
		IDs []string
	} `positional-args:"yes"`
	Run bool
}

func (rc *RemoveCmd) Execute(args []string) error {
	rc.Run = true
	return nil
}

type ClearCmd struct {
	Pos struct {
		Paths []string
	} `positional-args:"yes"`
	Run bool
}

func (cc *ClearCmd) Execute(args []string) error {
	cc.Run = true
	return nil
}

type EditCmd struct {
	Action string
	Text   string
	Editor bool `short:"e" long:"editor" description:"Use an editor instead of text entered as an argument."`
	Run    bool
}

func (ec *EditCmd) Execute(args []string) error {
	if len(args) > 0 {
		ec.Run = true
		ec.Action = args[0]
		if len(args) > 1 {
			ec.Text = args[1]
		}
	}
	return nil
}
