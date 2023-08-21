package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func startRepl() {

    scanner :=  bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("Please enter some text: ")

        scanner.Scan()
        
        text := scanner.Text()
        cleaned := cleanInput(text)

        // Checking for Enter input from CLI 
        if len(cleaned) == 0 {
            continue 
        }

        commandName := cleaned[0]
        availableCommands := getCommands() 

        command, ok := availableCommands[commandName]
        if !ok {
            fmt.Println("Invalid Command")
            continue
        }
        command.callback()
        switch command.name {
        case "help":
            callbackHelp()
        case "exit":
            callbackExit()
            default: 
            fmt.Println("Invalid Command")
        }


    }
}

type cliCommand struct {
    name string 
    description string 
    callback func() error
}

func getCommands() map[string]cliCommand {
    return map[string]cliCommand {
        "help": {
            name: "help",
            description: "Display all Options found in the Help Menu",
            callback: callbackHelp,
        },
        "exit": {
            name: "exit",
            description: "Stops the Pokedex CLI",
            callback: callbackExit, 
        },
    }
}


func cleanInput(str string) []string {
    lower_case := strings.ToLower(str)
    words := strings.Fields(lower_case)
    return words

}


