package main

import "fmt"

func recursividad() /*main*/ {

  // Iniciar la recursion con el número 100

  imprimirNumeros(100)

}

// Función recursiva para imprimir números del 100 al 0

func imprimirNumeros(n int) {
	
	//Cuando una función es recursiva, se repite a sí misma con un valor diferente de su parámetro hasta alcanzar una condición de parada (caso base).
	// Condición de parada: si n es menor que 0, termina la recursion

	if n < 0 {
			return
	}
	
	// Imprimir el número actual

	fmt.Println(n)

	// Llamada recursiva con el siguiente número menor

	imprimirNumeros(n - 1)

}

