package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nGEDCOM tools.")
		fmt.Fprintf(os.Stderr, "\n\nFlags:\n")
		flag.PrintDefaults()
	}

	var gedfile = flag.String("file", "", "GEDCOM filename")
	var ignorelist = flag.String("ignore", "", "rules ignore list")
	var verbose = flag.Bool("verbose", false, "verbose mode")
	var startPersonId = flag.String("id", "", "person id")

	flag.Parse()

	if *gedfile == "" {
		flag.Usage()
		os.Exit(1)
	}

	// printErrors()

	if *personId == "" {
		flag.Usage()
		os.Exit(1)
	}

	// getSocialTree(personId *string)

	// drawTimenet()
}
