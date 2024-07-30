package cli_cmd

import (
	"fmt"
	"os"
)

// type cliCommand struct {
// 	name		string
// 	description	string
// 	callback	func() error 
// }

type Cmds interface {
	CommandHello()
	CommandQuit()
}

func CommandHello() {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("<Command>  <Description>")
	fmt.Println("Help: Displays all commands and a brief description of what they do ")
	fmt.Println("Quit: Safely Exits the Program")
}

func CommandQuit() {
	os.Exit(0)
}
