package main

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type pokemonEncounter struct {
	Pokemon struct {
		Name string `json:"name"`
	} `json:"pokemon"`
}

type exploreResponse struct {
	PokemonEncounters []pokemonEncounter `json:"pokemon_encounters"`
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("please provide a location area name")
	}

	areaName := url.PathEscape(args[0])

	apiURL := fmt.Sprintf(
		"https://pokeapi.co/api/v2/location-area/%s/",
		areaName,
	)

	body, err := fetchAPIData(apiURL, cfg.cache)
	if err != nil {
		return err
	}

	area := exploreResponse{}
	if err := json.Unmarshal(body, &area); err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", args[0])

	if len(area.PokemonEncounters) == 0 {
		fmt.Println("No Pokémon found in this area.")
		return nil
	}

	fmt.Println("Found Pokémon:")

	for _, encounter := range area.PokemonEncounters {
		fmt.Println(" -", encounter.Pokemon.Name)
	}

	return nil
}
