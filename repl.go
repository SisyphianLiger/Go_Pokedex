package main

import (
	"bufio"
	"fmt"
	"github.com/SisyphianLiger/Go_Pokedex/cli_cmd"
	"os"
	"strings"
)

func repl() {
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
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
