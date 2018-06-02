package main

import "fmt"

// RemoveCmd tool command.
type RemoveCmd struct {
	Args struct {
		IDs []uint64 `positional-arg-name:"ACTION" description:"ID of action to remove."`
	} `positional-args:"yes"`
}

// Execute remove.
func (rc *RemoveCmd) Execute(args []string) error {
	if len(rc.Args.IDs) == 0 {
		pr("You need to specify one or more action IDs to remove.")
		return nil
	}

	act := loadCurrentOrFail()
	x := true
	for _, id := range rc.Args.IDs {
		_, ok := act.List[id]
		if ok {
			delete(act.List, id)
			if x == false {
				fmt.Printf(", %d", id)
			} else {
				fmt.Printf("Removed #%d", id)
				x = false
			}
		} else {
			pr("\nUnknown ID %d", id)
		}
	}
	pr("")
	return act.Save()
}
