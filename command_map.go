package main

import (
    "fmt"
    "log"

	"github.com/SisyphianLiger/Go_Pokedex/internal_pck/pokeapi"
)


func callbackMap() error {
	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	// Make it Pretty Output
	fmt.Print("Location Areas: \n")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
    return nil
}
