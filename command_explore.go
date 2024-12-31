package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	data, err := cfg.pokeapiClient.GetPokemanLists(args[0])
	if err != nil {
		return nil
	}

	fmt.Printf("Exploring %s\n", args[0])
	fmt.Printf("Found Pokemon\n")

	for _, Encounter := range data.PokemonEncounters {
		fmt.Printf("- %s\n", Encounter.Pokemon.Name)
	}

	return nil
}
