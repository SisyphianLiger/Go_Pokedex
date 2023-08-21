package pokeapi

type Poke_Mapper struct {
    Count       int          `json:"count"`
    Next_20     *string      `json:"name"`
    Previous    *string      `json:"previous"`
    Results     []struct {
        Name    string      `json:"name"`
        URL     string      `json:"url"`
    }`json:"results"`
}
