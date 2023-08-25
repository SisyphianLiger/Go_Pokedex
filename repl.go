package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func startRepl(cfg * config) {

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
        args := []string{}
        if len(cleaned) > 1 {
            args = cleaned[1:]
        }
        availableCommands := getCommands() 

        command, ok := availableCommands[commandName]
        if !ok {
            fmt.Println("Invalid Command")
            continue
        }
        
        command.callback(cfg, args...)
        
    }
}

type cliCommand struct {
    name string 
    description string 
    callback func(* config, ...string) error
}

func getCommands() map[string]cliCommand {
    return map[string]cliCommand {
        "h": {
            name: "help",
            description: "Display all Options found in the Help Menu",
            callback: callbackHelp,
        },
        "c": {
            name: "catch",
            description: "Catches a pokemon",
            callback: callbackCatch,
        },
        "i": {
            name: "inspect",
            description: "Inspects Pokemon",
            callback: callbackInspect,
        },
        "q": {
            name: "exit",
            description: "Stops the Pokedex CLI",
            callback: callbackExit, 
        },
        "m": {
            name: "map",
            description: "Shows Details of Pokemon Locations",
            callback: callbackMap, 
        },
        "p": {
            name: "pokedex",
            description: "Shows all names of captured pokemon",
            callback: callbackPokedex, 
        },
        "e": {
            name: "explore",
            description: "explore pokemons {location_area}",
            callback: callbackExp, 
        },
        "mb": {
            name: "map back",
            description: "Shows Details of Pokemon Locations",
            callback: callbackMb,
        },
        "r": {
            name: "relase",
            description: "relases captured pokemon",
            callback: callbackRelease,
        },
    }
}


func cleanInput(str string) []string {
    lower_case := strings.ToLower(str)
    words := strings.Fields(lower_case)
    return words

}


