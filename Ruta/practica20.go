package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

/* DIFICULTAD EXTRA:
Utilizando la PokeAPI (https://pokeapi.co), crea un programa por
terminal al que le puedas solicitar información de un Pokémon concreto
utilizando su nombre o número.
- Muestra el nombre, id, peso, altura y tipo(s) del Pokémon
- Muestra el nombre de su cadena de evoluciones
- Muestra los juegos en los que aparece
- Controla posibles errores
*/

// Estructuras para mapear los datos JSON
type Pokemon struct {
	Name       string        `json:"name"`
	ID         int           `json:"id"`
	Weight     int           `json:"weight"`
	Height     int           `json:"height"`
	Types      []PokemonType `json:"types"`
	GameIndices []GameIndex   `json:"game_indices"`
}

type PokemonType struct {
	Type TypeDetail `json:"type"`
}

type TypeDetail struct {
	Name string `json:"name"`
}

type Species struct {
	EvolutionChain EvolutionChain `json:"evolution_chain"`
}

type EvolutionChain struct {
	URL string `json:"url"`
}

type Evolution struct {
	Chain Chain `json:"chain"`
}

type Chain struct {
	Species    SpeciesDetail `json:"species"`
	EvolvesTo  []Chain       `json:"evolves_to"`
}

type SpeciesDetail struct {
	Name string `json:"name"`
}

type GameIndex struct {
	Version VersionDetail `json:"version"`
}

type VersionDetail struct {
	Name string `json:"name"`
}

func practica20() {
	const baseURL = "https://pokeapi.co/api/v2/pokemon/"
	var pokemonName string

	// Verificar si se proporciona un nombre de Pokémon como argumento en la línea de comandos
	if len(os.Args) < 2 {
		fmt.Println("Por favor, ingresa el nombre o número del Pokémon:")
		fmt.Scanln(&pokemonName)
	} else {
		pokemonName = os.Args[1]
	}

	pokemonName = strings.ToLower(strings.Replace(pokemonName, " ", "-", -1))

	// Obtener información del Pokémon
	pokemon, err := getPokemonInfo(baseURL + pokemonName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Obtener información de la especie para la cadena de evolución
	speciesURL := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon-species/%d/", pokemon.ID)
	evolutionChainURL, err := getEvolutionChainURL(speciesURL)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Obtener la cadena de evoluciones
	evolutionNames, err := getEvolutionChain(evolutionChainURL)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Obtener los juegos en los que aparece el Pokémon
	games := getGames(pokemon)

	// Mostrar la información del Pokémon
	fmt.Println("Nombre:", strings.Title(pokemon.Name))
	fmt.Println("ID:", pokemon.ID)
	fmt.Println("Peso:", pokemon.Weight)
	fmt.Println("Altura:", pokemon.Height)
	fmt.Println("Tipos:")
	for _, t := range pokemon.Types {
		fmt.Println("-", strings.Title(t.Type.Name))
	}
	fmt.Println("Evoluciones:")
	for _, name := range evolutionNames {
		fmt.Println("-", strings.Title(name))
	}
	fmt.Println("Aparece en los juegos:")
	for _, game := range games {
		fmt.Println("-", strings.Title(game))
	}
}

func getPokemonInfo(url string) (Pokemon, error) {
	resp, err := http.Get(url)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error realizando la petición: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Pokemon{}, fmt.Errorf("error: estado de respuesta %d", resp.StatusCode)
	}

	var pokemon Pokemon
	err = json.NewDecoder(resp.Body).Decode(&pokemon)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error decodificando JSON: %v", err)
	}

	return pokemon, nil
}

func getEvolutionChainURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error realizando la petición: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error: estado de respuesta %d", resp.StatusCode)
	}

	var species Species
	err = json.NewDecoder(resp.Body).Decode(&species)
	if err != nil {
		return "", fmt.Errorf("error decodificando JSON: %v", err)
	}

	return species.EvolutionChain.URL, nil
}

func getEvolutionChain(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error realizando la petición: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: estado de respuesta %d", resp.StatusCode)
	}

	var evolution Evolution
	err = json.NewDecoder(resp.Body).Decode(&evolution)
	if err != nil {
		return nil, fmt.Errorf("error decodificando JSON: %v", err)
	}

	var evolutionNames []string
	evolutionNames = append(evolutionNames, evolution.Chain.Species.Name)
	evolutionNames = append(evolutionNames, getEvolvesToNames(evolution.Chain.EvolvesTo)...)

	return evolutionNames, nil
}

func getEvolvesToNames(chains []Chain) []string {
	var names []string
	for _, chain := range chains {
		names = append(names, chain.Species.Name)
		names = append(names, getEvolvesToNames(chain.EvolvesTo)...)
	}
	return names
}

func getGames(pokemon Pokemon) []string {
	gameSet := make(map[string]struct{})
	for _, gameIndex := range pokemon.GameIndices {
		gameSet[gameIndex.Version.Name] = struct{}{}
	}

	var games []string
	for game := range gameSet {
		games = append(games, game)
	}
	return games
}