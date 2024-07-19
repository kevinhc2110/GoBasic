package main

import (
	"fmt"
	"os"
	"strings"
)

//* Problema 2 Substitution

func cifradoSustitucion() {
	// Verifica si se proporciona un único argumento de línea de comandos
	if len(os.Args) != 2 {
		fmt.Println("Error: El programa debe recibir exactamente un argumento.")
		os.Exit(1)
	}

	// Obtiene la clave de sustitución del argumento de línea de comandos
	clave := os.Args[1]

	// Valida la clave
	if !validarClave(clave) {
		fmt.Println("Error: La clave no es válida.")
		os.Exit(1)
	}

	// Solicita al usuario ingresar el texto sin formato
	fmt.Print("Texto plano: ")
	var textoSinFormato string
	fmt.Scan(&textoSinFormato)

	// Cifra el texto sin formato utilizando la clave de sustitución
	textoCifrado := cifrar(textoSinFormato, clave)

	// Imprime el texto cifrado
	fmt.Println("Texto cifrado:", textoCifrado)
}

// Función para validar la clave

func validarClave(clave string) bool {
	if len(clave) != 26 {
		return false // La clave debe tener exactamente 26 caracteres
	}

	// Creamos un mapa para almacenar las letras de la clave y su aparición

	letras := make(map[rune]bool)
	for _, char := range clave {
		// Verificamos si cada carácter es alfabético (mayúscula o minúscula)
		// El operador ! antes de la expresión invierte su valor booleano. Por lo tanto, esta condición verifica si char no es una letra minúscula y no es una letra mayúscula. Es decir, verifica si char no es una letra alfabética.
		// Es verdadero si char no es una letra alfabética.
		if !('a' <= char && char <= 'z') && !('A' <= char && char <= 'Z') {
			return false // La clave contiene caracteres no alfabéticos
		}
		// Verifica si cada letra aparece exactamente una vez
		if letras[char] {
			return false // La clave contiene letras duplicadas
		}
		// Registramos la aparición de la letra en el mapa
		letras[char] = true
	}
	// Si pasa todas las verificaciones, la clave es válida
	return true
}

// Función para cifrar el texto sin formato

func cifrar(textoSinFormato string, clave string) string {

	textoCifrado := ""

	// Iteramos sobre cada carácter en el texto sin formato
	for _, char := range textoSinFormato {
		// Verificamos si el carácter es una letra minúscula
		if 'a' <= char && char <= 'z' {
			// Concatenamos el carácter correspondiente en la clave, convertido a minúscula
			textoCifrado += strings.ToLower(string(clave[char-'a']))
			// Verificamos si el carácter es una letra mayúscula
		} else if 'A' <= char && char <= 'Z' {
			// Concatenamos el carácter correspondiente en la clave, convertido a mayúscula
			textoCifrado += strings.ToUpper(string(clave[char-'A']))
		} else {
			textoCifrado += string(char) // Conserva caracteres no alfabéticos
		}
	}
	return textoCifrado
}
