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

const baseURL = "https://pokeapi.co/api/v2/pokemon/"

type Pokemon struct {
	Name       string        `json:"name"`
	ID         int           `json:"id"`
	Weight     int           `json:"weight"`
	Height     int           `json:"height"`
	Types      []PokemonType `json:"types"`
	EvolvesTo  string        `json:"evolves_to"`
	AppearIn   []GameVersion `json:"appear_in"`
}

type PokemonType struct {
	Type TypeDetail `json:"type"`
}

type TypeDetail struct {
	Name string `json:"name"`
}

type GameVersion struct {
	Name string `json:"name"`
}

func main() {

	// Verificamos si al menos un argumento fue pasado en la línea de comandos

	if len(os.Args) < 2 {
		fmt.Println("Por favor, ingresa el nombre o número del Pokémon.")
		return
	}

	pokemonName := os.Args[1]

	// Reemplazamos espacios en blanco con guiones bajos

	pokemonName = strings.Replace(pokemonName, " ", "-", -1)

	// Realizamos la solicitud GET a la PokéAPI

	resp, err := http.Get(baseURL + pokemonName)
	if err != nil {
		fmt.Printf("Error realizando la petición: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Verificamos si la solicitud fue exitosa (código de estado 200)

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: estado de respuesta %d\n", resp.StatusCode)
		return
	}

	// Decodificamos el JSON de la respuesta en la estructura Pokemon

	var pokemon Pokemon
	err = json.NewDecoder(resp.Body).Decode(&pokemon)
	if err != nil {
		fmt.Printf("Error decodificando JSON: %v\n", err)
		return
	}

	// Mostramos la información del Pokémon

	fmt.Println("Nombre:", pokemon.Name)
	fmt.Println("ID:", pokemon.ID)
	fmt.Println("Peso:", pokemon.Weight)
	fmt.Println("Altura:", pokemon.Height)
	fmt.Println("Tipos:")

	for _, t := range pokemon.Types {
		fmt.Println("-", t.Type.Name)
	}
	fmt.Println("Evoluciona a:", pokemon.EvolvesTo)
	fmt.Println("Aparece en los juegos:")
	for _, g := range pokemon.AppearIn {
		fmt.Println("-", g.Name)
	}
}
