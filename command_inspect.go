package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	if len(args) != 1 {
		return errors.New("you must provide the name of the pokemon to inspect")
	}

	name := args[0]

	pokemon, exists := cfg.caughtPokemon[name]
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	// Imprimir la información del Pokémon
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}
