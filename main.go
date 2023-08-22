package main

import (
	"github.com/SisyphianLiger/Go_Pokedex/internal_pck/pokeapi"
)

// holding stateful info for callback
type config struct { 
    pokeapiClient pokeapi.Client
    nextLocationAreaURL *string
    prevLocationAreaURL *string
}


func main()  {

    cfg := config {
        pokeapiClient: pokeapi.NewClient(),
    }

    startRepl(&cfg) 
}
