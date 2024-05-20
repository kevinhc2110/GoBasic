package main

import "fmt"

func practica05() /*main*/ {
	x := 10
	y := 20

	 // Llamar a la función de intercambio por valor

	nuevoX, nuevoY := porValor(x, y)

	fmt.Println("Originales: x =", x, "y =", y) 
	fmt.Println("Intercambiados: nuevoX =", nuevoX, "nuevoY =", nuevoY)

	a := 30
  b := 40

	// Llamar a la función de intercambio por referencia 

	porReferencia(&a, &b)

	fmt.Println("Intercambiados por referencia: a =", a, "b =", b) 
}


func porValor(valor1 int, valor2 int) (int, int) {
	return valor1, valor2
}

// Se usan puntero para acceder a la dirección de memoria y cambiar el valor cuando lo necesitemos

func porReferencia(valor1 *int, valor2 *int) {

	*valor1, *valor2 = *valor2, *valor1
}