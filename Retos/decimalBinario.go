package main

import "fmt"

/*
Crea un programa se encargue de transformar un número
decimal a binario sin utilizar funciones propias del lenguaje que lo hagan directamente.
*/

func decimalBinario() {

	var numero int

	fmt.Println("Ingrese una palabra par invertir")
	fmt.Scanln(&numero)

	binario := convertirEntero(numero)
	fmt.Println(binario)

}

func convertirEntero(numero int) []int {

	binario := []int{}

	// Convertir el número decimal a binario
	// Ciclo while

	for numero > 0 {
		residuo := numero % 2
		binario = append(binario, residuo)
		numero /= 2
	}

	for i, j := 0, len(binario)-1; i < j; i, j = i+1, j-1 {
		binario[i], binario[j] = binario[j], binario[i]
	}

	return binario

}
