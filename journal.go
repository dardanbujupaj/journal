package main

import (
	"flag"
	"fmt"
	"strings"
	"github.com/dardanbujupaj/journal/entry"
)

func main() {
    // how many entries are printed (-n)
    var nEntries int
    flag.IntVar(&nEntries, "n", 5, "Amount of Entries to print")
    // how many entries are printed (-a)
    var allEntries bool
    flag.BoolVar(&allEntries, "a", false, "Print all available entries")

	flag.Parse()

	args := flag.Args()
	
	var command string

    // default command -> list entries
	if (len(args) == 0) {
		command = "list"
	} else {
		command = args[0]
	}
		
	switch (command) {
		case "list", "l", "ls":
            var amount int = nEntries
            if (allEntries) {
                amount = -1
            }
            // list entries
			entry.ListEntries(amount)
		case "create", "c", "cr":
            // create new journal entry with given text
			entry.CreateEntry(strings.Join(args[1:], " "))

		default:
            // print help text
			fmt.Println("not a valid command: " + command)
            flag.PrintDefaults()
	}
}

