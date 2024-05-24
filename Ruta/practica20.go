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

const baseURL = "https://pokeapi.co/api/v2/pokemon/" // URL base de la PokéAPI

// Definición de estructuras para mapear los datos JSON
type Pokemon struct {
	Name   string        `json:"name"`
	ID     int           `json:"id"`
	Weight int           `json:"weight"`
	Height int           `json:"height"`
	Types  []PokemonType `json:"types"`
}

type PokemonType struct {
	Type TypeDetail `json:"type"`
}

type TypeDetail struct {
	Name string `json:"name"`
}

func main() {
	
	var pokemonName string

	// Verificar si se proporciona un nombre de Pokémon como argumento en la línea de comandos

	if len(os.Args) < 2 {
		// Si no se proporciona, solicitar al usuario que ingrese el nombre o número del Pokémon
		fmt.Println("Por favor, ingresa el nombre o número del Pokémon:")
		fmt.Scanln(&pokemonName) // Leer la entrada del usuario desde la terminal
	} else {
		pokemonName = os.Args[1] // Usar el argumento proporcionado como el nombre del Pokémon
	}

	// Formatear el nombre del Pokémon para que coincida con el formato de la URL de la PokéAPI

	pokemonName = strings.Replace(pokemonName, " ", "-", -1)

	// Realizar la solicitud GET a la PokéAPI para obtener información del Pokémon

	resp, err := http.Get(baseURL + pokemonName)
	if err != nil {
		fmt.Printf("Error realizando la petición: %v\n", err)
		return
	}
	defer resp.Body.Close() // Cerrar el cuerpo de la respuesta después de que la función main haya terminado

	// Verificar si la solicitud fue exitosa (código de estado 200)

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: estado de respuesta %d\n", resp.StatusCode)
		return
	}

	// Decodificar el JSON de la respuesta en la estructura Pokemon

	var pokemon Pokemon
	err = json.NewDecoder(resp.Body).Decode(&pokemon)
	if err != nil {
		fmt.Printf("Error decodificando JSON: %v\n", err)
		return
	}

	// Mostrar la información del Pokémon

	fmt.Println("Nombre:", pokemon.Name)
	fmt.Println("ID:", pokemon.ID)
	fmt.Println("Peso:", pokemon.Weight)
	fmt.Println("Altura:", pokemon.Height)
	fmt.Println("Tipos:")

	// Mostrar los tipos del Pokémon

	for _, t := range pokemon.Types {
		fmt.Println("-", t.Type.Name)
	}
}
