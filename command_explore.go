package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("What do you want to explore? Please add as argument")
	}
	pokemonResp, err := cfg.pokeapiClient.ExploreLocation(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", args[0])
	fmt.Println("Found Pokemon:")
	for _, loc := range pokemonResp.PokemonEncounters {
		fmt.Printf(" - %s\n", loc.Pokemon.Name)
	}
	return nil
}
