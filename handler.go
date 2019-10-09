package main

import "fmt"

const (
	baseLeaguepediaURL = "https://lol.gamepedia.com/"
	version            = "1.0.0"
)

// Prints the list of accepted commands
func printHelp() {
	fmt.Printf("LoL Pro Player Stats (version %s)\n", version)
	fmt.Println("Available commands:")
	fmt.Println("* -h | --help\t Prints the list of available commands")
	fmt.Println("* -v | --version Prints the version of the application")
}

// Prints the current version of the application
func printVersion() {
	fmt.Printf("LoL Pro Player Stats version %s\n", version)
}

/* Handles a request to lookup information related to a player
 * Receives:
 * args ([]string) - Arguments passed in the terminal by the user
 */
func handlePlayerOptions(args []string) {
	var (
		queryURL, leagueCode, split, week string
	)
}
