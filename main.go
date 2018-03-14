// act
// A TODO list for the command line.

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/asdine/storm"
	"github.com/jessevdk/go-flags"
)

type Path struct {
	ID   int    `storm:"id"`
	Name string `storm:"index,unique"`
}

type Action struct {
	ID     int    `storm:"id,increment"`
	Text   string `storm:"index,unique"`
	PathID int    `storm:"index"`
}

var opts struct {
	Ver    VersionCmd `command:"version" description:"Print program version and exit." alias:"ver"`
	All    AllCmd     `command:"all" description:"List all actions for all paths."`
	Add    AddCmd     `command:"add" description:"Add a new action." alias:"a" alias:"new"`
	Remove RemoveCmd  `command:"remove" description:"Remove one or more actions." alias:"rem" alias:"rm" alias:"delete" alias:"del"`
	Clear  ClearCmd   `command:"clear" description:"Remove a path and its contents." alias:"clr"`
	Edit   EditCmd    `command:"change" description:"Change the text of an entry." alias:"edit"`
}

var db storm.DB

func main() {

	db, err := storm.Open(defaultDB())
	check(err)
	defer db.Close()

	pathname, err := os.Getwd()
	check(err)
	pathname, err = filepath.Abs(pathname)
	check(err)

	// Simplest case: Show all actions for current path
	if len(os.Args) < 2 {
		var path Path
		err = db.One("Name", pathname, &path)
		if err == storm.ErrNotFound {
			pr("No entries for current path.")
			return
		}
		var act []Action
		err = db.Find("PathID", path.ID, &act)
		if err == storm.ErrNotFound || len(act) == 0 {
			pr("No entries for current path.")
			return
		}
		printActions(&act)
		return
	}

	_, err = flags.Parse(&opts)
	if err != nil {
		return
	}

	if opts.Ver.Run {
		pr("%s %s", program, Version)
		return
	}

	// Add a new entry for the current path
	if opts.Add.Run {
		if opts.Add.Action == "" {
			pr("You need to specify an action string to add.")
			return
		}
		var path Path
		err = db.One("Name", pathname, &path)
		if err == storm.ErrNotFound {
			path = Path{Name: pathname}
			db.Save(&path)
		}

		var act Action
		err = db.One("Text", opts.Add.Action, &act)
		if err == storm.ErrNotFound {
			act = Action{Text: opts.Add.Action, PathID: path.ID}
			db.Save(&act)
			pr("Added entry #%d.", act.ID)
		}
		return
	}

	// Show EVERYTHING
	if opts.All.Run {
		var act []Action
		err = db.AllByIndex("ID", &act)
		if err == storm.ErrNotFound {
			pr("There are no entries for any path.")
			return
		}
		printAllActions(db, &act)
		return
	}

	// Remove actions
	if opts.Remove.Run {
		if len(opts.Remove.Pos.IDs) == 0 {
			pr("You need to specify one or more action IDs to remove.")
			return
		}
		act := Action{}
		for _, a := range opts.Remove.Pos.IDs {
			act.ID, _ = strconv.Atoi(a)
			err = db.DeleteStruct(&act)
			if err == nil {
				pr("Removed #%s.", a)
			} else {
				pr("Unknown ID %s.", a)
			}
		}
		return
	}

	// Clear paths
	if opts.Clear.Run {
		for _, pathname := range opts.Clear.Pos.Paths {
			pathname, _ := filepath.Abs(pathname)
			var path Path
			err = db.One("Name", pathname, &path)
			if err != storm.ErrNotFound {
				pr("Clearing %s", pathname)
				var act []Action
				err = db.Find("PathID", path.ID, &act)
				if err == nil {
					for _, a := range act {
						db.DeleteStruct(&a)
					}
				}
			} else {
				pr("Unknown path %s.", pathname)
			}
		}
		return
	}

	// Edit actions
	if opts.Edit.Run {
		if opts.Edit.Editor {
			editAction(db, opts.Edit.Action)
		} else {
			if opts.Edit.Text != "" {
				n, _ := strconv.Atoi(opts.Edit.Action)
				var act Action
				err = db.One("ID", n, &act)
				check(err)
				act.Text = opts.Edit.Text
				err = db.Update(&act)
				check(err)
				pr("Changed #%s to %s", opts.Edit.Action, opts.Edit.Text)
			}
		}
		return
	}
}

func printActions(act *[]Action) {
	s := fmt.Sprintf("%d", len(*act))
	l := len(s)
	s = fmt.Sprintf("%d", l)
	for _, x := range *act {
		pr("%"+s+"d: %s", x.ID, x.Text)
	}
}

func printAllActions(sdb *storm.DB, act *[]Action) {
	s := fmt.Sprintf("%d", len(*act))
	l := len(s)
	s = fmt.Sprintf("%d", l)

	sections := map[string]int{}
	all := map[int][]Action{}
	var list []Action
	for _, a := range *act {
		list = all[a.PathID]
		if list == nil {
			list = []Action{}
			var path Path
			err := sdb.One("ID", a.PathID, &path)
			check(err)
			sections[path.Name] = path.ID
		}
		list = append(list, Action{a.ID, a.Text, a.PathID})
		all[a.PathID] = list
	}
	for k, v := range sections {
		pr("%s:", k)
		for _, a := range all[v] {
			pr("%"+s+"d: %s", a.ID, a.Text)
		}
		pr("")
	}
}
