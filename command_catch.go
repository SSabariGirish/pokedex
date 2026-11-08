package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/url"
	"strings"
)

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("please provide a Pokemon name")
	}

	pokemonName := strings.Join(args, "-")

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	apiURL := fmt.Sprintf(
		"https://pokeapi.co/api/v2/pokemon/%s/",
		url.PathEscape(pokemonName),
	)

	body, err := fetchAPIData(apiURL, cfg.cache)
	if err != nil {
		return err
	}

	pokemon := Pokemon{}
	if err := json.Unmarshal(body, &pokemon); err != nil {
		return err
	}

	catchChance := 100 - pokemon.BaseExperience/4

	if catchChance > 90 {
		catchChance = 90
	}

	if catchChance < 10 {
		catchChance = 10
	}

	if rand.Intn(100) < catchChance {
		cfg.pokedex[pokemon.Name] = pokemon

		fmt.Printf("%s was caught!\n", pokemon.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
