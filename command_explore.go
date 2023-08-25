package main

import (
	"errors"
	"fmt"
)




func callbackExp(cfg * config, args ...string) error {
    if len(args) != 1 {
        return errors.New("No Pokemon in the Area")
    }

    poke_loc_area := args[0]
 
    poke_loc, err := cfg.pokeapiClient.GetPokeLocation(poke_loc_area)
    if err != nil {
		return err
	}
	// Make it Pretty Output
	fmt.Printf("Pokemon in Area %s: \n", poke_loc.Name)
	for _, pokemon := range poke_loc.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name) 
	}
    return nil
}
