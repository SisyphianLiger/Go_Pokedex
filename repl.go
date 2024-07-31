package main

import (
	"bufio"
	"fmt"
	"github.com/SisyphianLiger/Go_Pokedex/cli_cmd"
	"os"
	"strings"
)

func run_repl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())

		cmdName := words[0]

		command, exists := cli_cmd.GetCommands()[cmdName]

		if exists {
			err := command.Callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			err := cli_cmd.CommandHelp()
			if err != nil {
				fmt.Println("Problem with the Program, Closing now")
				cli_cmd.CommandQuit()
			}
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
