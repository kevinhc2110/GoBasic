package main

import "fmt"

func fibonacci() {

	secFibonacci()
}

func secFibonacci(){
	
	n := 50 

	// Inicializamos los primeros dos términos de la secuencia

	a, b := 0, 1

	// Imprimimos el primer término

	fmt.Printf("%d ", a)

	// Iteramos para calcular e imprimir los siguientes 49 términos

	for i := 1; i < n; i++ {
		fmt.Printf("%d ", b)

		// Calculamos el siguiente término
		// La expresión a, b = b, a+b se evalúa en dos pasos:
		// Primero, se evalúan las expresiones a la derecha del operador =
		// b    -> 1  // El valor actual de `b`
		// a+b  -> 0+1 -> 1  // La suma de los valores actuales de `a` y `b`
		// Luego, se asignan estos valores a las variables a la izquierda del operador =

		a, b = b, a+b
	}
	fmt.Println() // Nueva línea al final de la secuencia		
}

