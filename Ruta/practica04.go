//* Cadena de caracteres

package main

/*
DIFICULTAD EXTRA:
Crea un programa que analice dos palabras diferentes y realice comprobaciones
para descubrir si son:
- Palíndromos
- Anagramas
- Isogramas
*/

import (
	"fmt"
	"strings"
)

func practica04() /*main*/ {

	var palabra1, palabra2 string

	fmt.Println("Ingrese la primera palabra:")
	fmt.Scanln(&palabra1)
	fmt.Println("Ingrese la segunda palabra:")
	fmt.Scanln(&palabra2)

	if esPalindromo(palabra1) {
		fmt.Printf("%s es un palíndromo.\n", palabra1)
		} else {
			fmt.Printf("%s no es un palíndromo.\n", palabra1)
	}

	if esPalindromo(palabra2) {
		fmt.Printf("%s es un palíndromo.\n", palabra2)
		} else {
			fmt.Printf("%s no es un palíndromo.\n", palabra2)
	}


	if esAnagrama(palabra1, palabra2) {
		fmt.Printf("%s y %s son anagramas.\n", palabra1, palabra2)
	} else {
		fmt.Printf("%s y %s no son anagramas.\n", palabra1, palabra2)
	}

	if esIsograma(palabra1) {
		fmt.Printf("%s es un isograma.\n", palabra1)
		} else {
			fmt.Printf("%s no es un isograma.\n", palabra1)
	}

	if esIsograma(palabra2) {
		fmt.Printf("%s es un isograma.\n", palabra2)
		} else {
			fmt.Printf("%s no es un isograma.\n", palabra2)
	}

}

func esPalindromo (palabra string) bool {

	palabra= strings.ToLower(palabra)

	//for i := 0; i < len(palabra)/2; i++ {: Inicia un bucle for que recorre la mitad de la longitud de la palabra. Esto es así porque queremos comparar los caracteres desde el principio de la palabra con los caracteres desde el final de la palabra.
	//if palabra[i] != palabra[len(palabra)-1-i] { return false }: En cada iteración del bucle, compara el carácter en la posición i con el carácter en la posición len(palabra)-1-i, que es el carácter correspondiente desde el final de la palabra.

	for i := 0; i < len(palabra)/2; i++ {
		if palabra[i] != palabra[len(palabra)-1-i] {
			return false
		}
	}
	return true
}

func esAnagrama(palabra1 string, palabra2 string) bool {

	// Este es un bucle for que itera sobre cada carácter (letra) de la cadena palabra.
	// El range palabra produce dos valores en cada iteración: el índice del carácter y el carácter en sí. En este caso, el índice no se utiliza, por lo que se descarta con _.
	// letra toma el valor de cada carácter en palabra en cada iteración del bucle.
	// mapeoPalabra2 es un mapa (map) en Go que tiene caracteres (rune) como claves y enteros (int) como valores.
	// El mapa mapeoPalabra2 se utiliza para contar la frecuencia de cada carácter en palabra.
	// mapeoPalabra2[letra] accede al valor en el mapa asociado con la clave letra.
	//	El operador ++ incrementa el valor asociado con esa clave en 1.
	//	Si la clave letra no existía previamente en el mapa, Go inicializa automáticamente el valor a 0 antes de aplicar el incremento.

	palabra1 = strings.ToLower(palabra1)
	palabra2 = strings.ToLower(palabra2)

	if len(palabra1) != len(palabra2){
		return false
	}

	mapeoPalabra1 := make(map[rune]int)
	mapeoPalabra2 := make(map[rune]int)

	for _, letra := range palabra1{
		mapeoPalabra1[letra]++
	}

	for _, letra := range palabra1{
		mapeoPalabra2[letra]++
	}

	for letra, frecuencia := range mapeoPalabra1 {
		if mapeoPalabra2[letra] != frecuencia {
			return false
		}
	}

	return true
}

func esIsograma(palabra string) bool {

	// if frecuencia[letra] > 0 { return false }: Dentro del bucle, verificamos si la letra actual ya está presente en el mapa frecuencia.
	// Si la frecuencia de la letra es mayor que 0, significa que ya se ha encontrado previamente en la palabra, lo que indica que la palabra no es un isograma. En este caso, la función retorna false.

	palabra = strings.ToLower(palabra)
	frecuencia := make(map[rune]int)

	for _, letra := range palabra {
		if frecuencia[letra] > 0 {
			return false
		}
		frecuencia[letra]++
	}

	return true
}







