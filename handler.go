package main

import (
	"fmt"
	"os"

	"github.com/anaskhan96/soup"
)

const (
	baseLeaguepediaURL = "https://lol.gamepedia.com"
	version            = "1.0.0"
)

// Prints the list of accepted commands
func printHelp() {
	fmt.Printf("LoL Pro Player Stats (version %s)\n", version)
	fmt.Println("Available commands:")
	fmt.Println("* -h | --help\t Prints the list of available commands")
	fmt.Println("* -v | --version Prints the version of the application")
	fmt.Println("* PLAYER CODE SPLIT YEAR Prints the statistics of a player in a given split of a given year (e.g. lol-pro-player-stats.exe Perkz LEC Summer 2019)")
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
		queryURL, player, leagueCode, split, year string
	)

	player, leagueCode, split, year = args[1], args[2], args[3], args[4]
	queryURL = baseLeaguepediaURL + "/Special:RunQuery/MatchHistoryPlayer?MHP%5Bpreload%5D=Player&MHP%5Btournament%5D=" + leagueCode + "%20" + year + "%20" + split + "&MHP%5Bspl%5D=yes&MHP%5Blink%5D=" + player + "&pfRunQueryFormName=MatchHistoryPlayer"
	resp, err := soup.Get(queryURL)
	if err != nil {
		handleError(err)
	}

	printPlayerTable(player, leagueCode, split, year, resp)
}

/* Print a player's champion stats in a given split of a given year.
 * Receives:
 * player (string) - Player's in game handle
 * leagueCode (string) - Code of the league (LCS, LEC, etc.)
 * split (string) - Split of the league season (Spring, Summer, Winter)
 * year (string) - Year of the season
 * document (string) - Leaguepedia page HTML document
 */
func printPlayerTable(player, leagueCode, split, year, document string) {
	fmt.Printf(" %s %s %s %s Stats\n\n", player, leagueCode, split, year)
	fmt.Println(" Champion \t|  G  |  W  |  L  | WR    | KDA")

	soup.SetDebug(true)
	doc := soup.HTMLParse(document)
	championTable := doc.Find("table", "class", "spstats")
	championTableBody := championTable.Find("tbody")
	championTableBodyRows := championTableBody.Children()

	for _, row := range championTableBodyRows {
		// only want TD elements inside tbody
		if row.Children()[0].NodeValue == "td" {
			championName := row.Children()[0].Children()[1].NodeValue
			championNameDistance := calculateStringDistance(15, championName)

			games := row.Children()[1].Children()[0].Text()
			gamesDistance := calculateStringDistance(1, games)

			wins := row.Children()[2].Text()
			losses := row.Children()[3].Text()

			winratio := row.Children()[4].Text()
			winratioDistance := calculateStringDistance(4, winratio)
			kda := row.Children()[8].Text()

			fmt.Printf("%s %s|  %s %s|  %s  |  %s  | %s %s| %s\n", championName, championNameDistance, games, gamesDistance, wins, losses, winratio, winratioDistance, kda)
		}
	}
}

func calculateStringDistance(baseDistance int, name string) string {
	stringDistanceLen := baseDistance - len(name)
	stringDistance := ""
	for i := 0; i <= stringDistanceLen; i++ {
		stringDistance = stringDistance + " "
	}

	return stringDistance
}

func handleError(err error) {
	fmt.Printf("%s", err)
	os.Exit(1)
}
