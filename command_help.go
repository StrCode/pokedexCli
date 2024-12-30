package main

import "fmt"

func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	commands := getCommands()

	for key, value := range commands {
		fmt.Printf("%s: %s\n", key, value.description)
	}

	return nil
}
