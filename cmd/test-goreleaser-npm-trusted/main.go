package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	showVersion := flag.Bool("version", false, "print version and exit")
	flag.Parse()

	if *showVersion {
		fmt.Printf("test-goreleaser-npm-trusted %s (commit %s, built at %s)\n", version, commit, date)
		return
	}

	fmt.Println("Hello from test-goreleaser-npm-trusted!")
	if len(os.Args) > 1 {
		fmt.Printf("args: %v\n", os.Args[1:])
	}
}
