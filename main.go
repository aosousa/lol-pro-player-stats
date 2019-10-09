package main

import "os"

func main() {
	args := os.Args

	if len(args) == 1 {
		cmd := args[1]

		switch cmd {
		case "-h", "--help":
			printHelp()
		case "-v", "--version":
			printVersion()
		default:
			printHelp()
		}
	} else if len(args) > 1 {
		handlePlayerOptions(args)
	} else {
		printHelp()
	}
}
