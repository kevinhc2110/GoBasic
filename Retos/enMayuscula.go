package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
Crea una función que reciba un String de cualquier tipo y se encargue de
poner en mayúscula la primera letra de cada palabra.
- No se pueden utilizar operaciones del lenguaje que
  lo resuelvan directamente.
*/

func esMayuscula() {

	// Opción 1

	str := "hola mundo, qué tal estás"

	result := ponerMayus(str)
	fmt.Println("Resultado:", result)

	// Opción 2

	result2 := ponerMayus2(str)
	fmt.Println("Resultado:", result2)

}

// Opción 1

func ponerMayus(str string) string {

	// Convertir la cadena de entrada a un slice de runas
	runes := []rune(str)

	// Indicador para rastrear si el siguiente carácter debe ser mayúscula
	nextCharIsUpper := true

	// Iterar sobre cada carácter en la cadena
	for i, char := range runes {
		// Verificar si el carácter actual es una letra
		if unicode.IsLetter(char) {
			// Si el siguiente carácter debe ser mayúscula, convertir este carácter a mayúscula
			if nextCharIsUpper {
				runes[i] = unicode.ToUpper(char)
				// Actualizar el indicador para el siguiente carácter
				nextCharIsUpper = false
			}
		} else {
			// Si el carácter actual no es una letra, el siguiente carácter debe ser mayúscula
			nextCharIsUpper = true
		}
	}

	// Convertir el slice de runas de nuevo a una cadena y retornarla
	return string(runes)
}

// Opción 2

func ponerMayus2(cadena string) string {

	// Inicializar una cadena vacía para almacenar las palabras en mayúscula
	cadenaCapitalizada := ""

	// Dividir la cadena de entrada en palabras usando espacios como delimitadores
	palabras := strings.Fields(cadena)

	// Iterar sobre cada palabra
	for _, palabra := range palabras {

		// Capitalizar la primera letra de la palabra actual
		// Notación de segmentación de cadenas de Go para extraer una subcadena que comienza en el índice 0 (el primer carácter) y tiene una longitud de 1 carácter

		primeraLetra := strings.ToUpper(palabra[:1])

		// La expresión palabra[1:] extrae una subcadena de palabra que comienza en el índice 1 (es decir, el segundo carácter) y se extiende hasta el final de la cadena.
		// Esta subcadena contiene todas las letras excepto la primera, que se ha almacenado en primeraLetra.

		restoPalabra := palabra[1:]

		// Añadir la palabra en mayúscula a la cadena capitalizada
		cadenaCapitalizada += primeraLetra + restoPalabra + " "
	}

	// Eliminar el espacio en blanco final
	cadenaCapitalizada = strings.TrimSpace(cadenaCapitalizada)

	// Devolver la cadena capitalizada
	return cadenaCapitalizada
}
