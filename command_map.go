package main

import (
	"fmt"
	"log"
)

func callbackMap(cfg * config) error {

    resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		log.Fatal(err)
	}
    fmt.Println("%v", cfg.pokeapiClient)
    fmt.Println("%v", resp)
	// Make it Pretty Output
	fmt.Print("Location Areas: \n")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name) 
	}
    cfg.nextLocationAreaURL = resp.Next
    cfg.prevLocationAreaURL = resp.Previous
    return nil
}
