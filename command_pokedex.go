package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, pokeman := range cfg.caughtPokemon {
		fmt.Printf("  - %s\n", pokeman.Name)
	}
	return nil
}
