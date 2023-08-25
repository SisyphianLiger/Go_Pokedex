package main

import (
	"errors"
	"fmt"	
    "math/rand"
)


func callbackCatch(cfg * config, args ...string) error {
    if len(args) != 1 {
        return errors.New("Not a location area")
    }

    pokemon_name:= args[0]
 
    pokemon, err := cfg.pokeapiClient.GetPokemon(pokemon_name)
    if err != nil {
		return err
	}

    // Making a Catch Function
    const threshold = 20
    rand_num := rand.Intn(pokemon.BaseExperience)

    if rand_num < threshold {
        fmt.Printf("Failed to Catch %s\n", pokemon.Name)
        return fmt.Errorf("Failed to Catch %s", pokemon.Name)
    }

	// Make it Pretty Output
	fmt.Printf("%s been caught!\n", pokemon.Name)
    cfg.caughtPokemon[pokemon_name] = pokemon
    return nil
}
