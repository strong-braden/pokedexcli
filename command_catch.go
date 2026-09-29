package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("What pokemon are you catching? Please add a Pokemon as an argument")
	}
	pokemonResp, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}
	exp := pokemonResp.BaseExperience
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonResp.Name)
	if rand.Intn(255) < exp {
		fmt.Printf("%s escaped!\n", pokemonResp.Name)
	} else {
		fmt.Printf("%s was caught!\n", pokemonResp.Name)
		cfg.caughtPokemon[pokemonResp.Name] = pokemonResp
	}
	// for _, loc := range pokemonResp.PokemonEncounters {
	// 	fmt.Printf(" - %s\n", loc.Pokemon.Name)
	// }
	return nil
}
