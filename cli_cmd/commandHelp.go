package cli_cmd

import "fmt"

func CommandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("-----------------------")
	fmt.Println()
	fmt.Println("<Command>  <Description>")
	fmt.Println("Help: Displays all commands and a brief description of what they do ")
	fmt.Println("Quit: Safely Exits the Program")
	fmt.Println("-----------------------")
	fmt.Println()
	return nil
}
