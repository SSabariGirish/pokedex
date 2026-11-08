package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type locationResponse struct {
	Count    int
	Next     string
	Previous string
	Results  []location
}

var PokeBaseURL = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"

func fetchAPIData(url string, cache *Cache) ([]byte, error) {
	// Check cache first
	if data, ok := cache.Get(url); ok {
		//fmt.Println("[CACHE HIT]", url)
		return data, nil
	}

	//fmt.Println("[NETWORK REQUEST]", url)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if !json.Valid(body) {
		return nil, fmt.Errorf("invalid JSON response")
	}

	// Cache only a successful response containing valid JSON
	cache.Add(url, body)

	return body, nil
}

func fetchLocationAreas(url string, cache *Cache) (locationResponse, error) {
	body, err := fetchAPIData(url, cache)
	if err != nil {
		return locationResponse{}, err
	}

	area := locationResponse{}
	if err := json.Unmarshal(body, &area); err != nil {
		return locationResponse{}, err
	}

	return area, nil
}

func commandMap(cfg *config, args ...string) error {

	cmd := cfg.commands["map"]

	url := PokeBaseURL
	if cmd.nextURL != "" {
		url = cmd.nextURL
	}

	area, err := fetchLocationAreas(url, cfg.cache)
	if err != nil {
		return err
	}

	for _, loc := range area.Results {
		fmt.Println(loc.Name)
	}

	cmd.previousURL = area.Previous
	cmd.nextURL = area.Next

	cfg.commands["map"] = cmd

	return nil
}

func commandMapB(cfg *config, args ...string) error {

	cmd := cfg.commands["map"]

	if cmd.previousURL == "" {
		fmt.Println("You're on the first page.")
		return nil
	}

	area, err := fetchLocationAreas(cmd.previousURL, cfg.cache)
	if err != nil {
		return err
	}

	for _, loc := range area.Results {
		fmt.Println(loc.Name)
	}

	cmd.previousURL = area.Previous
	cmd.nextURL = area.Next

	cfg.commands["map"] = cmd

	return nil
}
