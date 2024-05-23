package main

import "fmt"

// En Go, el concepto de callback se puede implementar utilizando funciones como valores, lo que permite pasar funciones como argumentos a otras funciones.
// Esto ilustra el concepto de callback en Go, donde una función puede aceptar otras funciones como argumentos y llamarlas según sea necesario.

// Definimos un tipo de función llamado Callback que toma un entero y devuelve un entero.

type Callback func(int) int

// La función procesarNumeros toma una lista de números y una función de callback,
// y aplica la función de callback a cada número de la lista.

func procesarNumeros(numeros []int, callback Callback) {
	for _, num := range numeros {
		resultado := callback(num)
		fmt.Println("Resultado:", resultado)
	}
}

// Función de ejemplo que duplica un número.

func duplicar(numero int) int {
	return numero * 2
}

// Función de ejemplo que suma 10 a un número.

func sumar10(numero int) int {
	return numero + 10
}

func callbacks() /*main*/ {
	
	// Creamos una lista de números.

	numeros := []int{1, 2, 3, 4, 5}

	fmt.Println("Aplicando la función duplicar:")
	
	// Llamamos a procesarNumeros con la función duplicar como callback.

	procesarNumeros(numeros, duplicar)

	fmt.Println("\nAplicando la función sumar10:")
	
	// Llamamos a procesarNumeros con la función sumar10 como callback.

	procesarNumeros(numeros, sumar10)
}