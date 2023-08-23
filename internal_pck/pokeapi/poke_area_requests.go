package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (Poke_Mapper, error) {
    endpoint := "/location-area"
    fullURL := baseURL + endpoint
    // Help rotate between map 
    if pageURL != nil {
        fullURL = *pageURL
    }
     

    request, error := http.NewRequest("GET", fullURL, nil)

    if error != nil {
        return Poke_Mapper{}, error
    }
    
    // processes the request
    response, err := c.httpClient.Do(request)

    if error != nil {
        return Poke_Mapper{}, err
    }

    defer response.Body.Close()

    if response.StatusCode > 399 {
        return Poke_Mapper{}, fmt.Errorf("Bad Status Cade")
    }

    // Check the Data
    data, err := io.ReadAll(response.Body)
    if err != nil {
        return Poke_Mapper{}, err 
    }

    // Finally we unpack the data here
    PokeMapper := Poke_Mapper{}
    err = json.Unmarshal(data, &PokeMapper)

    if err != nil {
        return Poke_Mapper{}, err 
    }

    return PokeMapper, nil
}
