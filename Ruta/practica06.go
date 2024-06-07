package main

import "fmt"

/*
DIFICULTAD EXTRA:
Utiliza el concepto de recursividad para:
- Calcular el factorial de un número concreto (la función recibe ese número).
- Calcular el valor de un elemento concreto (según su posición) en la
sucesión de Fibonacci (la función recibe la posición).
*/

func practica06()/*main*/ {

	var numero int;
	
	fmt.Println("Ingrese un numero")
	fmt.Scanln(&numero)
	
	numeroFactorial := factorial(numero)

	fmt.Println("El factorial de", numero, "es:", numeroFactorial)

	var posicion int;

	fmt.Println("Ingrese una posición Fibonacci")
	fmt.Scanln(&posicion)

	enesimoFibonacci := fibonacci(posicion)
	
	fmt.Println("El número Fibonacci en la posición", posicion, "es", enesimoFibonacci)
	
}


func factorial(n int)int {

	// Caso base: Si n es 0, el factorial de 0 es 1

	if n == 0 {
		return 1
	}

	// Caso recursivo: n multiplicado por el factorial de (n-1)
	return n * factorial((n-1))
}

func fibonacci(n int) int {
	if n <= 1 {
			return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}