package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("What pokemon are you inspecting? Please add a Pokemon as an argument")
	}
	pokemonResp, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}
	pokemon, ok := cfg.caughtPokemon[pokemonResp.Name]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}
	fmt.Println("Name: " + pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, slot := range pokemon.Types {
		fmt.Printf("  - %s\n", slot.Type.Name)
	}
	// for _, loc := range pokemonResp.PokemonEncounters {
	// 	fmt.Printf(" - %s\n", loc.Pokemon.Name)
	// }
	return nil
}
