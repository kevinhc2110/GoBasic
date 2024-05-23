package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
Escribe una función que reciba dos palabras (String) y retorne
verdadero o falso (Bool) según sean o no anagramas.
- Un Anagrama consiste en formar una palabra reordenando TODAS
las letras de otra palabra inicial.
- NO hace falta comprobar que ambas palabras existan.
- Dos palabras exactamente iguales no son anagrama.
*/

func Anagrama() {
	
	if esAnagrama1("amor", "roma") {
		fmt.Println("Las palabras son anagramas")
	} else {
		fmt.Println("Las palabras no son anagramas")
	}

	if esAnagrama2("amor", "roma") {
		fmt.Println("Las palabras son anagramas")
	} else {
		fmt.Println("Las palabras no son anagramas")
	}
}

// Opción 1

func esAnagrama1(palabra1, palabra2 string) bool {

	palabra1 = strings.ToLower(strings.ReplaceAll(palabra1, " ", ""))
	palabra2 = strings.ToLower(strings.ReplaceAll(palabra2, " ", ""))

	if len(palabra1) != len(palabra2) {
		return false
	}

	runas1 := []rune(palabra1)
	runas2 := []rune(palabra2)

	// sort.Slice() es una función del paquete sort de Go que ordena un slice utilizando un algoritmo de ordenación personalizado.
	// La función de comparación func(i, j int) bool {return runas1[i] < runas1[j]} compara dos runas en los índices i y j del slice runas1. Devuelve true si la runa en el índice i es menor que la runa en el índice j, lo que significa que la runa en el índice i debe estar antes en la secuencia ordenada.

	sort.Slice(runas1, func(i, j int) bool { return runas1[i] < runas1[j] })
	sort.Slice(runas2, func(i, j int) bool { return runas2[i] < runas2[j] })

	return string(runas1) == string(runas2)
}

// Opción 2

func esAnagrama2(palabra1, palabra2 string) bool {
	
	if len(palabra1) != len(palabra2) {
			return false
	}

	//Recorre cada carácter (c) en la palabra palabra1.
  //Si el carácter c ya existe como clave en el mapa cuenta1, su valor (número de veces que aparece) se incrementa en 1.
	
	cuenta1 := make(map[rune]int)
	for _, c := range palabra1 {
			cuenta1[c]++
	}

	cuenta2 := make(map[rune]int)
	for _, c := range palabra2 {
			cuenta2[c]++
	}

	// Recorre cada elemento (clave-valor) del mapa cuenta1
	// La variable c toma el valor de la clave (la letra).
	//La variable n toma el valor asociado (número de veces que aparece la letra).

	for c, n := range cuenta1 {

		// Esta línea verifica si la frecuencia de la letra c en el mapa cuenta2 (palabra2) coincide con la frecuencia en cuenta1 (palabra1).
		
		if cuenta2[c] != n {
				return false
		}
	}

	return true
}