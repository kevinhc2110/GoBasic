package main

import "fmt"

/*
Crea un programa que invierta el orden de una cadena de texto
sin usar funciones propias del lenguaje que lo hagan de forma automática.
- Si le pasamos "Hola mundo" nos retornaría "odnum aloH"
*/

func invirtiendoCadenas() {

	var palabra string

	fmt.Println("Ingrese una palabra par invertir")
	fmt.Scanln(&palabra)

	fmt.Println(invertirString(palabra))
}

func invertirString(palabra string) string {

	runas := []rune(palabra)

	// El len siempre esta un numero por encima pues las matrices empiezan desde 0

	for i, j := 0, len(runas)-1; i < j; i, j = i+1, j-1 {

		// El valor de runas[j] = 3 (que es 'a') se asigna a runas[i] = 0.
		// El valor de runas[i] = 0 (que es 'h') se asigna a runas [j] = 3

		runas[i], runas[j] = runas[j], runas[i]
	}
	return string(runas)
}
