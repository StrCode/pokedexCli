package main

import (
	"time"

	"github.com/StrCode/pokedexCli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		caughtPokemon: make(map[string]pokeapi.Pokemon),
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
