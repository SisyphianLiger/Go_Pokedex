package main

import (
	"errors"
	"fmt"	
)

func callbackPokedex(cfg * config, args ...string) error {
    fmt.Printf("Current Pokemon from Pokedex\n")
    for _, pokemon := range cfg.caughtPokemon {
        fmt.Printf("%s\n", pokemon.Name)
    }
    return nil 
}

func callbackInspect(cfg * config, args ...string) error {
    if len(args) != 1 {
        return errors.New("Only one pokemon at a time")
    }

    pokemon_name:= args[0]
    poke_res, ok := cfg.caughtPokemon[pokemon_name]
 
    if !ok {
		return fmt.Errorf("No pokemon with name %s found", pokemon_name) 
	}
    
    fmt.Printf("Name: %s\n", poke_res.Name)
    fmt.Printf("Height: %v\n", poke_res.Height)
    fmt.Printf("Weight: %v\n", poke_res.Weight)
    fmt.Printf("Stats:")
    for _, stat := range poke_res.Stats {
        fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
    }
    fmt.Printf("Types:")
    for _, t := range poke_res.Types {
        fmt.Printf("Type: %v\n", t.Type.Name)
    }
	// Make it Pretty Output
    return nil
}
