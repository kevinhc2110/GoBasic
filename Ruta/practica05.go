//* Valor y referencia

package main

import "fmt"

/*
DIFICULTAD EXTRA:
Crea dos programas que reciban dos parámetros (cada uno) definidos como
variables anteriormente.
- Cada programa recibe, en un caso, dos parámetros por valor, y en otro caso, por referencia.
Estos parámetros los intercambia entre ellos en su interior, los retorna, y su retorno
se asigna a dos variables diferentes a las originales. A continuación, imprime
el valor de las variables originales y las nuevas, comprobando que se ha invertido
su valor en las segundas.
Comprueba también que se ha conservado el valor original en las primeras.
*/

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