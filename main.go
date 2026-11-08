package main

import "time"

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Get Command Help",
		callback:    commandHelp,
	},
	"map": {
		name:        "map",
		description: "Visit 20 areas",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Visit 20 areas back",
		callback:    commandMapB,
	},
	"explore": {
		name:        "explore",
		description: "Explore a location area for Pokémon",
		callback:    commandExplore,
	},
	"catch": {
		name:        "catch",
		description: "Catch a Pokemon",
		callback:    commandCatch,
	},
	"inspect": {
		name:        "inspect",
		description: "Inspect a caught Pokemon",
		callback:    commandInspect,
	},
	"pokedex": {
		name:        "pokedex",
		description: "List all caught Pokemon",
		callback:    commandPokedex,
	},
}

func main() {
	cfg := &config{
		commands: commands,
		cache:    NewCache(5 * time.Minute),
		pokedex:  make(map[string]Pokemon),
	}

	startRepl(cfg)
}
