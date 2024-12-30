package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			"help",
			"Displays a help message",
			commandHelp,
		},
		"exit": {
			"exit",
			"Exit the Pokedex",
			commandExit,
		},
		"map": {
			"map",
			"Display the locations",
			CommandMap,
		},
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}

func startRepl() {
	// Scanner waits for user input

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("Pokedex > ")

		for !scanner.Scan() {
			break
		}

		input := scanner.Text()

		if len(input) == 0 {
			continue
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}

		words := cleanInput(input)
		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}
