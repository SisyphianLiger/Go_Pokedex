package pokeapi

import (
	"net/http"
	"time"

	"github.com/SisyphianLiger/Go_Pokedex/internal_pck/pokecache"
)

// Base url
const baseURL = "https://pokeapi.co/api/v2"


type Client struct {
    httpClient http.Client 
    cache pokecache.Cache
}


func NewClient(interval time.Duration) Client {
    return Client {
        cache: pokecache.NewCache(interval),
        httpClient: http.Client {
            Timeout: time.Minute,
        },
    }
}


