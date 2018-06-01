package main

var opts struct {
	Ver    VersionCmd `command:"version" description:"Print program version and exit." alias:"ver"`
	All    AllCmd     `command:"all" description:"List all actions for all paths."`
	Add    AddCmd     `command:"add" description:"Add a new action." alias:"a" alias:"new"`
	Remove RemoveCmd  `command:"remove" description:"Remove one or more actions." alias:"rem" alias:"rm" alias:"delete" alias:"del"`
	Clear  ClearCmd   `command:"clear" description:"Remove a path and its contents." alias:"clr"`
	Edit   EditCmd    `command:"change" description:"Change the text of an entry." alias:"edit"`
}
