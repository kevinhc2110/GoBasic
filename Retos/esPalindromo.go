package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

/*
Escribe una función que reciba un texto y retorne verdadero o
falso (Boolean) según sean o no palíndromos.
Un Palíndromo es una palabra o expresión que es igual si se lee
de izquierda a derecha que de derecha a izquierda.
NO se tienen en cuenta los espacios, signos de puntuación y tildes.
Ejemplo: Ana lleva al oso la avellana.
*/

func palindromo() {

	fmt.Println("Ingresa la cadena para analizar: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	texto := scanner.Text()

	if esPalindromo(texto) {
		fmt.Printf("El texto \"%s\" es un palíndromo.\n", texto)
	} else {
		fmt.Printf("El texto \"%s\" no es un palíndromo.\n", texto)
	}
	
}

func eliminarTildes(texto string) string {

	sinTildes := map[rune]rune{
		'á': 'a', 'é': 'e', 'í': 'i', 'ó': 'o', 'ú': 'u',
		'Á': 'A', 'É': 'E', 'Í': 'I', 'Ó': 'O', 'Ú': 'U',
		'ñ': 'n', 'Ñ': 'N',
		'ü': 'u', 'Ü': 'U',
	}

	var textoSinTildes strings.Builder

	for _, char := range texto {

		// cambio, encontrado := sinTildes[char]: Intenta encontrar el carácter char en el mapa 
		// sinTildes. Si se encuentra, cambio contendrá el valor asociado (el carácter sin tilde y encontrado será true
		// Si no se encuentra, cambio será el valor cero del tipo rune y encontrado será false.
		if cambio, encontrado := sinTildes[char]; encontrado{
			textoSinTildes.WriteRune(cambio)
		}else {
			textoSinTildes.WriteRune(char)
		}
	}

	return textoSinTildes.String()
}

func esPalindromo(texto string) bool {

	texto = strings.ToLower(texto)
	texto = eliminarTildes(texto)

	var filtrado strings.Builder

	for _, char := range texto {
		if unicode.IsLetter(char){
			filtrado.WriteRune(char)
		}
	}

	textoFiltrado := filtrado.String()

	reverso := reverse(textoFiltrado)


	return textoFiltrado == reverso
}

func reverse(texto string)string  {

	runas := []rune(texto)

	for i,j := 0, len(runas)-1; i<j ; i,j = i+1, j-1 {
		runas[i], runas[j] = runas[j], runas[i]
	}
	
	return string(runas)
}