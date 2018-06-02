# act [![Build status](https://travis-ci.org/Urethramancer/act.svg)](https://travis-ci.org/Urethramancer/act)
This is a command-line tool for TODO lists. Each action you add is sorted under whichever working path you're currently in.

##  Dependencies
1. Go (tested on version 1.10)
2. [cross](https://github.com/Urethramancer/cross)
3. [go-flags](https://github.com/jessevdk/go-flags)

## Platforms
Tested on Linux and macOS. Prepared for various BSD-derivatives. Not yet ready for Windows.

## Usage
Running the program without arguments or commands lists all actions registered for the current path:

```sh
$ act
133: Push this to GitHub if all looks well.
135: Give the README a once-over.
```

Using the `all` command shows everything for all paths:

```sh
$ act all
/Users/orb/src/go/act:
133: Push this to GitHub if all looks well.
135: Give the README a once-over.

/Users/orb/src/go/todos:
 22: Scrap this and make something better.
```

Adding a new action is straight-forward:

```sh
$ act add "Stop making useless TODO entries."
Added entry #142.
```

You can change an entry for the current path with the `change` command, or its alias `edit`:

```sh
$ act change 135 "Let somebody else eyeball the README."
Changed #135 to "Let somebody else eyeball the README."
```

If you specify the `-e` flag instead of text the current version of the entry specified will open in an editor specified in the **EDITOR** environment variable instead:

```sh
$ act change 135 -e
```

Entries for the current path can be removed by with `remove` or its aliases, `rem`, `delete` or `del`:

```sh
$ act rem 142
Removed #142.
```

It's also easy to clear all entries for a path with `clear` or `clr`:

```sh
$ act clr /Users/orb/src/go/todos
Clearing /Users/orb/src/go/todos
```

And of course you can pipe output into a file to build a TODO list you can check into version control:

```sh
$ act > TODO
```

## Licence
MIT.
