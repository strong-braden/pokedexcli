package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// ExploreLocation -
func (c *Client) ExploreLocation(name string) (RespLocationArea, error) {
	url := baseURL + "/location-area/" + name

	if val, ok := c.cache.Get(url); ok {
		areaResp := RespLocationArea{}
		err := json.Unmarshal(val, &areaResp)
		if err != nil {
			return RespLocationArea{}, err
		}

		return areaResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationArea{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return RespLocationArea{}, errors.New("location data for " + name + " not found")
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationArea{}, err
	}

	areaResp := RespLocationArea{}
	err = json.Unmarshal(dat, &areaResp)
	if err != nil {
		return RespLocationArea{}, err
	}

	c.cache.Add(url, dat)
	return areaResp, nil
}
