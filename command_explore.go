package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
    if len(args) != 1 {
        return errors.New("you must provide a location name")
    }
	
    locationArea := args[0]

    fmt.Printf("Exploring %s...\n", locationArea)

    areaResp, err := cfg.pokeapiClient.ExploreLocationArea(locationArea)
    if err != nil {
        return err
    }

    fmt.Println("Found Pokemon:")
    for _, p := range areaResp.PokemonEncounters {
        fmt.Printf(" - %s\n", p.Pokemon.Name)
    }

    return nil
}

