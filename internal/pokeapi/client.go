package pokeapi

import (
	"net/http"
	"time"
    "encoding/json"
    "fmt"
    "io"
	"github.com/Mariomck1209/pokedexcli/internal/pokecache"
)

// Client -
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

type LocationArea struct {
    PokemonEncounters []struct {
        Pokemon struct {
            Name string `json:"name"`
        } `json:"pokemon"`
    } `json:"pokemon_encounters"`
}

func (c *Client) ExploreLocationArea(areaName string) (LocationArea, error) {
    // Revisa si ya está en caché
    if data, ok := c.cache.Get(areaName); ok {
        var area LocationArea
        err := json.Unmarshal(data, &area)
        if err != nil {
            return LocationArea{}, err
        }
        return area, nil
    }

    // Si no está, haz request
    url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", areaName)

    resp, err := c.httpClient.Get(url)
    if err != nil {
        return LocationArea{}, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return LocationArea{}, err
    }

    var area LocationArea
    err = json.Unmarshal(body, &area)
    if err != nil {
        return LocationArea{}, err
    }

    // Guarda en caché
    c.cache.Add(areaName, body)

    return area, nil
}
