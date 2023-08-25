package main


import "fmt"


func callbackHelp(cfg * config, args ...string) error {

    fmt.Println("Welcome to the Pokedex help menu:")
    fmt.Println("***********Commands**************")
    availableCommands := getCommands()
    for _, command := range availableCommands {
        fmt.Printf("Command: (%v) %s: %s\n",string(command.name[0]), command.name, command.description)
    }
    fmt.Println("")
    return nil
}
