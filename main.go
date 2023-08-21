package main

import "github.com/SisyphianLiger/Go_Pokedex/internal_pck/pokeapi"

// holding stateful info for callback
type config struct { 
    pokeapiClient pokeapi.Client
    Poke_Map_next *string
    Poke_Map_prev *string
}


func main()  {

    cfg := config {
        pokeapiClient: pokeapi.NewClient(),
    }

    startRepl(&cfg) 
}
