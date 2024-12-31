package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("enter the name of the pokemon")
	}

	name := args[0]

	pokemonResp, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	random := rand.Intn(pokemonResp.BaseExperience)

	if random < 40 {
		fmt.Printf("%s escaped\n", name)
	} else {
		fmt.Printf("%s was caught\n", name)
		cfg.caughtPokemon[name] = pokemonResp
	}

	return nil
}
