package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func contandoPalabras() {

	// Opción 1

	fmt.Print("Ingrese el nombre del contacto a buscar:  ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	cadena := scanner.Text()

	// Contar las palabras en la frase
	conteo := conteoPalabras(cadena)

	// Mostrar el recuento de palabras
	fmt.Println("\nRecuento de palabras:")
	for palabra, contador := range conteo {
		fmt.Printf("%s: %d\n", palabra, contador)
	}

	// Opción 2

	wordCounts := make(map[string]int) // Crea un mapa para almacenar el recuento de palabras

	fmt.Println("Introduce el texto:") 
	scanner1 := bufio.NewScanner(os.Stdin) 
	for scanner1.Scan() { // Itera sobre cada línea de entrada
			line := scanner1.Text() // Obtiene la línea de entrada como una cadena de texto
			words := strings.Fields(line) // Divide la línea en palabras utilizando los espacios como delimitadores
			for _, word := range words { // Itera sobre cada palabra
					cleanedWord := cleanWord(word) // Limpia la palabra de caracteres no alfabéticos y la convierte en minúsculas
					wordCounts[cleanedWord]++ // Incrementa el recuento de la palabra en el mapa
			}
	}

	// Imprimir recuento final de palabras
	fmt.Println("\nRecuento final de palabras:") 
	for word, count := range wordCounts { // Itera sobre el mapa de recuento de palabras
			fmt.Printf("%s: %d\n", word, count) // Imprime la palabra y su recuento
	}

}

// Opción 1

func conteoPalabras(cadena string) map[string]int {

	// Inicializar un mapa para contar las palabras
	conteo := make(map[string]int)

	// Eliminar signos de puntuación y convertir la frase a minúsculas
	cadena = limpiarCadena(cadena)

	// Dividir la frase en palabras
	palabras := strings.Fields(cadena)

	// Contar las palabras
	for _, palabra := range palabras {
		conteo[palabra]++
	}

	return conteo
}

func limpiarCadena(cadena string) string {
	
	// Convertir la cadena a minúsculas
	cadena = strings.ToLower(cadena)

	// Definir los signos de puntuación que queremos eliminar
	signosPuntuacion := "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

	// Eliminar los signos de puntuación de la cadena
	for _, puntuacion := range signosPuntuacion {
		cadena = strings.ReplaceAll(cadena, string(puntuacion), "")
	}

	return cadena
}

// Opción 2

// Función cleanWord limpia una palabra de caracteres no alfabéticos y la convierte en minúsculas

func cleanWord(word string) string {
	var cleanedWord strings.Builder // Declara un strings.Builder para almacenar la palabra limpia
	for _, char := range word { // Itera sobre cada carácter de la palabra
			if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') { // Verifica si el carácter es alfabético
					cleanedWord.WriteRune(char) // Agrega el carácter al strings.Builder
			}
	}
	return strings.ToLower(cleanedWord.String()) // Devuelve la palabra limpia y en minúsculas
}
