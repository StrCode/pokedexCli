package pokeapi

import (
	"net/http"
	"time"

	"github.com/StrCode/pokedexCli/internal/pokecache"
)

// Client -
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

// NewClient -
func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		pokecache.NewCache(cacheInterval),
		http.Client{
			Timeout: timeout,
		},
	}
}
