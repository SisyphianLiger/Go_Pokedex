package pokeapi

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

func (c *Client) ListLocationAreas() (Poke_Mapper, error) {
    endpoint := "/location-area"
    fullURL := baseURL + endpoint

    request, error := http.NewRequest("GET", fullURL, nil)

    if error != nil {
        return Poke_Mapper{}, error
    }

    response, err := c.httpClient.Do(request)
    if error != nil {
        return Poke_Mapper{}, err
    }

    if response.StatusCode > 399 {
        return Poke_Mapper{}, fmt.Errorf("Bad Status Cade")
    }

    data, err := io.ReadAll(response.Body)

    if err != nil {
        return Poke_Mapper{}, err 
    }

    PokeMapper := Poke_Mapper{}
    err = json.Unmarshal(data, &PokeMapper)

    if err != nil {
        return Poke_Mapper{}, err 
    }

    return PokeMapper, nil
}
