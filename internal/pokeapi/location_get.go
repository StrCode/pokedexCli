package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, nil
	}

	if cache, err := c.cache.Get(url); err {
		locationResp := Location{}

		if err := json.Unmarshal(cache, &locationResp); err != nil {
			return Location{}, nil
		}

		return locationResp, nil

	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, nil
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, nil
	}

	locationResp := Location{}

	if err := json.Unmarshal(dat, &locationResp); err != nil {
		return Location{}, nil
	}

	c.cache.Add(url, dat)

	return locationResp, nil
}
