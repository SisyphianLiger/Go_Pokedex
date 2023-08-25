package main

import (
	"errors"
	"fmt"
)

func callbackRelease(cfg * config, args ...string) error {
    if len(args) != 1 {
        return errors.New("Only one pokemon at a time")
    }

    pokemon_rel := args[0]

    ok := cfg.caughtPokemon[pokemon_rel] 
    if ok.Name != pokemon_rel {
        return fmt.Errorf("%s not in list", ok.Name)
    }
 
    fmt.Printf("Good bye %s\n", pokemon_rel)
    delete(cfg.caughtPokemon, pokemon_rel)

    return nil 
}

