package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg * config, args ...string) error {

    resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}
	// Make it Pretty Output
	fmt.Print("Location Areas: \n")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name) 
	}
    cfg.nextLocationAreaURL = resp.Next
    cfg.prevLocationAreaURL = resp.Previous
    return nil
}


func callbackMb(cfg * config, args ...string) error {

    if cfg.prevLocationAreaURL == nil {
        fmt.Println("You are on the First Page")
        return errors.New("You are on the First Page")
    }
    
    resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL)
	if err != nil {
		return err
	}
	// Make it Pretty Output
	fmt.Print("Location Areas: \n")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name) 
	}
    cfg.nextLocationAreaURL = resp.Next
    cfg.prevLocationAreaURL = resp.Previous
    return nil
}
