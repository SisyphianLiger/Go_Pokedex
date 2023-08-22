package main


import "fmt"


func callbackHelp(cfg * config) error {

    fmt.Println("Welcome to the Pokedex help menu:")
    fmt.Println("***********Commands**************")
    availableCommands := getCommands()
    for _, command := range availableCommands {
        fmt.Printf("Command: %s: %s\n", command.name, command.description)
    }
    fmt.Println("")
    return nil
}
