package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


// GetLocationAreas makes a GET request to the PokeAPI location-area endpoint
// with the specified location area name and returns a LocationArea struct
func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	// build the endpoint URL
	endpoint := "/pokemon/" + pokemonName
	fullURL := baseURL + endpoint

	// check the cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		// cache hit
		pokemon := Pokemon{}

		// unmarshal the JSON data from the cache into the LocationAreasResp struct
		err := json.Unmarshal(dat, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}

		return pokemon, nil
	}

	// create a new http request with the specified method (GET), URL, and no body
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	// execute the request using the httpClient
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	// defer the closing of the response body until the function exits
	defer resp.Body.Close()

	// check the status code of the response
	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("request failed with status: %d", resp.StatusCode)
	}

	// read the response body into a byte array
	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	// create a new instance of the LocationAreasResp struct
	pokemon := Pokemon{}

	// unmarshal the JSON data from the response into the LocationAreasResp struct
	err = json.Unmarshal(dat, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	// add the URL to the cache
	c.cache.Add(fullURL, dat)

	// return the LocationAreasResp struct
	return pokemon, nil
}

