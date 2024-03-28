package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ListLocationAreas performs a GET request to the PokeAPI location-area endpoint
// and returns a LocationAreasResp struct
func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreasResp, error) {
	// build the endpoint URL
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if pageUrl != nil {
		fullURL = *pageUrl
	}

	// check the cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		// create a new instance of the LocationAreasResp struct
		locationAreaResp := LocationAreasResp{}

		// unmarshal the JSON data from the response into the LocationAreasResp struct
		err := json.Unmarshal(dat, &locationAreaResp)
		if err != nil {
			// if there was an error unmarshaling the data, return the error
			return LocationAreasResp{}, err
		}
		// return the LocationAreasResp struct
		return locationAreaResp, nil
	}
	fmt.Println("cache miss!")

	// create a new http request with the specified method (GET), URL, and no body
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		// if there was an error creating the request, return the error
		return LocationAreasResp{}, err
	}

	// execute the request using the httpClient
	resp, err := c.httpClient.Do(req)
	if err != nil {
		// if there was an error executing the request, return the error
		return LocationAreasResp{}, err
	}

	// defer the closing of the response body until the function exits
	// this ensures that the response body is always closed, even if we return early
	defer resp.Body.Close()

	// check the status code of the response
	// if it is greater than 399 (which is a client error or server error)
	if resp.StatusCode > 399 {
		// return an error with the status code as a string
		return LocationAreasResp{}, fmt.Errorf("request failed with status: %d", resp.StatusCode)
	}

	// read the response body into a byte array
	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		// if there was an error reading the response body, return the error
		return LocationAreasResp{}, err
	}

	// create a new instance of the LocationAreasResp struct
	locationAreaResp := LocationAreasResp{}

	// unmarshal the JSON data from the response into the LocationAreasResp struct
	err = json.Unmarshal(dat, &locationAreaResp)
	if err != nil {
		// if there was an error unmarshaling the data, return the error
		return LocationAreasResp{}, err
	}

	// add the URL to the cache
	c.cache.Add(fullURL, dat)

	// return the LocationAreasResp struct
	return locationAreaResp, nil

}

// GetLocationAreas makes a GET request to the PokeAPI location-area endpoint
// with the specified location area name and returns a LocationArea struct
func (c *Client) GetLocationAreas(locationAreaName string) (LocationArea, error) {
	// build the endpoint URL
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint

	// check the cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		// cache hit
		locationArea := LocationArea{}

		// unmarshal the JSON data from the cache into the LocationAreasResp struct
		err := json.Unmarshal(dat, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}

		return locationArea, nil
	}

	// create a new http request with the specified method (GET), URL, and no body
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	// execute the request using the httpClient
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}

	// defer the closing of the response body until the function exits
	defer resp.Body.Close()

	// check the status code of the response
	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("request failed with status: %d", resp.StatusCode)
	}

	// read the response body into a byte array
	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	// create a new instance of the LocationAreasResp struct
	locationArea := LocationArea{}

	// unmarshal the JSON data from the response into the LocationAreasResp struct
	err = json.Unmarshal(dat, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	// add the URL to the cache
	c.cache.Add(fullURL, dat)

	// return the LocationAreasResp struct
	return locationArea, nil
}

