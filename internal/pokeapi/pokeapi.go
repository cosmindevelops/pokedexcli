package pokeapi

import (
	"net/http"
	"time"

	"github.com/cosmindevelops/pokedexcli/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

// Create a new instance of the pokeapi client.
// Set http client timeout to 1 minute.
// This allows for customization of the http client.
func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
