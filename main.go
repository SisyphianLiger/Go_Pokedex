package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/SisyphianLiger/Go_Pokedex/cli_cmds"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Pokedex > ")
	for scanner.Scan() {
		switch strings.ToLower(scanner.Text()) {

		case "help":
			cli_cmd.CommandHello()

		case "quit":
			cli_cmd.CommandQuit()

		default:
			fmt.Printf("Select help or quit\n")
			fmt.Println("")
			fmt.Printf("Pokedex > ")
		}
	}
}
