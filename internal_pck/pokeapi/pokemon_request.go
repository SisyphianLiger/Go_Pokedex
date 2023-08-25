package pokeapi


import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemon_name string) (Pokemon, error) {
    endpoint := "/pokemon/" + pokemon_name
    fullURL := baseURL + endpoint
    
    dat, ok := c.cache.Get(fullURL)
    if ok {
        Pokes := Pokemon{}
        err := json.Unmarshal(dat, &Pokes)

        if err != nil {
            return Pokemon{}, err 
        }

        return Pokes, nil
    }

    request, error := http.NewRequest("GET", fullURL, nil)

    if error != nil {
        return Pokemon{}, error
    }
    
    // processes the request
    response, err := c.httpClient.Do(request)

    if error != nil {
        return Pokemon{}, err
    }

    defer response.Body.Close()

    if response.StatusCode > 399 {
        return Pokemon{}, fmt.Errorf("Bad Status Cade")
    }

    // Check the Data
    data, err := io.ReadAll(response.Body)
    if err != nil {
        return Pokemon{}, err 
    }

    // Finally we unpack the data here
    Pokes := Pokemon{}
    err = json.Unmarshal(data, &Pokes)

    if err != nil {
        return Pokemon{}, err 
    }
    
    c.cache.Add(fullURL, data)

    return Pokes, nil
}


